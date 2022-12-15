package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go.uber.org/zap"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
	"wash_bonus/intapi"
	"wash_bonus/internal/app"
	"wash_bonus/internal/app/balance"
	"wash_bonus/internal/app/wash_server"
	"wash_bonus/internal/dal"
	"wash_bonus/internal/firebase_authorization"
	"wash_bonus/internal/transport/grpc"
	"wash_bonus/internal/transport/rest"
	"wash_bonus/pkg/bootstrap"
)

func main() {
	cfg, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln("new config: ", err)
	}

	l, err := bootstrap.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalln("new logger: ", err)
	}

	dbConn, err := bootstrap.NewDbConn(cfg.DB)
	if err != nil {
		l.Fatalln("new db conn: ", err)
	}
	defer dbConn.Close()

	l.Debug("connected to db")

	err = bootstrap.UpMigrations(dbConn.DB, cfg.DB.Database, "migrations")
	if err != nil {
		l.Fatalln("up migrations: ", err)
	}

	l.Debug("applied migrations")

	authSvc := firebase_authorization.New(cfg.FirebaseConfig.FirebaseKeyFilePath)

	repo := dal.New(dbConn, l)

	userSvc := app.NewUserService(l, repo)
	balanceSvc := balance.New(l, repo)
	washServerSvc := wash_server.New(l, repo)

	intapi := grpc.New(l, balanceSvc, washServerSvc)

	errc := make(chan error)

	go runHTTPServer(errc, l, cfg, authSvc, userSvc, balanceSvc)
	go runGRPCServer(errc, l, cfg, intapi)

	err = <-errc
	if err != nil {
		l.Fatalln("rest api serve:", err)
	}

	l.Info("started server at: ", cfg.HTTPPort)
}

func runHTTPServer(errc chan error, l *zap.SugaredLogger, cfg *bootstrap.Config, authSvc firebase_authorization.Service, userSvc app.UserService, balanceSvc balance.Service) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("panic: ", r)
		}
	}()
	server, err := rest.NewServer(cfg, authSvc, l, userSvc, balanceSvc)
	if err != nil {
		l.Fatalln("init rest server:", err)
	}

	errc <- server.Serve()
}

func runGRPCServer(errc chan error, l *zap.SugaredLogger, cfg *bootstrap.Config, grpcSvc *grpc.Service) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("panic: ", r)
		}
	}()
	serverOptions := []grpc2.ServerOption{}

	if cfg.GrpcConfig.EnableTLS {
		credentialsTLS, err := loadTLSCredentials(cfg)
		if err != nil {
			errc <- fmt.Errorf("grpc: %v", err)
			return
		}
		serverOptions = append(serverOptions, grpc2.Creds(credentialsTLS))
	}

	server := grpc2.NewServer(serverOptions...)

	intapi.RegisterServerServiceServer(server, grpcSvc)
	intapi.RegisterSessionServiceServer(server, grpcSvc)

	listener, err := net.Listen("tcp", ":"+cfg.GrpcConfig.Port)
	if err != nil {
		errc <- fmt.Errorf("grpc: %v", err)
		return
	}

	errc <- server.Serve(listener)
}

func loadTLSCredentials(cfg *bootstrap.Config) (credentials.TransportCredentials, error) {
	pemClientCA, err := os.ReadFile(cfg.GrpcConfig.ClientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(cfg.GrpcConfig.ServerCertFile, cfg.GrpcConfig.ServerKeyFile)
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
		ClientCAs:    certPool,
	}), nil
}
