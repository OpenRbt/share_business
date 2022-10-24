package def

import (
	_ "github.com/joho/godotenv/autoload"
)

var (
	GRPCPort         = strGetenv("GRPC_PORT", "8091")
	GRPCEnableTLS    = boolGetenv("GRPC_ENABLE_TLS")
	ClientCACertFile = pathGetenv("GRPC_CLIENT_CA_CERT_FILE", "~/certs/ca-cert.pem")
	ServerCertFile   = pathGetenv("GRPC_SERVER_CERT_FILE", "~/certs/server-cert.pem")
	ServerKeyFile    = pathGetenv("GRPC_SERVER_KEY_FILE", "~/certs/server-key.pem")
)
