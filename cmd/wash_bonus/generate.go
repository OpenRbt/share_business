package main

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./intapi/grpc.proto
//go:generate rm -rf ./openapi/restapi
//go:generate swagger generate server -t ./openapi/ -f ./openapi/swagger.yaml --strict-responders --strict-additional-properties --principal wash_bonus/internal/app.Auth --exclude-main
//go:generate swagger generate client -t ./openapi/ -f ./openapi/swagger.yaml --strict-responders --strict-additional-properties --principal wash_bonus/internal/app.Auth
//go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;
