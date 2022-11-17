package sharebusiness

//go:generate rm -rf ./internal/transport/rest/restapi
//go:generate swagger generate server -t ./internal/transport/rest -f ./swagger.yaml --exclude-main
//go:generate swagger generate client -t ./internal/transport/rest -f ./swagger.yaml

//go:generate protoc --go_out=./internal/transport/grpc --go_opt=paths=source_relative --go-grpc_out=./internal/transport/grpc --go-grpc_opt=paths=source_relative ./internal/transport/grpc/grpc.proto
