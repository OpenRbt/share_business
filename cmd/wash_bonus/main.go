package main

import (
	"log"
	"wash_bonus/pkg/bootstrap"
)

//go:generate rm -rf ./openapi/restapi
//go:generate swagger generate server -t ./openapi/ -f ./openapi/swagger.yaml --strict-responders --strict-additional-properties --principal wash_bonus/internal/app.Auth --exclude-main
//go:generate swagger generate client -t ./openapi/ -f ./openapi/swagger.yaml --strict-responders --strict-additional-properties --principal wash_bonus/internal/app.Auth
//go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;

////go:generate protoc --go_out=./transport/grpc --go_opt=paths=source_relative --go-grpc_out=./transport/grpc --go-grpc_opt=paths=source_relative ./transport/grpc.proto

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

	//init repo

	//init app services

	//init rest

	l.Info("started server at: ", cfg.HTTPPort)
}
