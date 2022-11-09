package sharebusiness

//go:generate rm -rf openapi/models openapi/restapi openapi/client
//go:generate swagger generate server -t openapi/ -f swagger.yaml --exclude-main
//go:generate swagger generate client -t openapi/ -f swagger.yaml

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/transport/grpc/grpc.proto
