package sharebusiness

//go:generate rm -rf ./internal/api/restapi
//go:generate swagger generate server -t ./internal/transport/api/restapi -f ./swagger.yaml --exclude-main
//go:generate swagger generate client -t ./internal/transport/api/restapi -f ./swagger.yaml

//go:generate protoc --go_out=./internal/transport/grpc --go_opt=paths=source_relative --go-grpc_out=./internal/transport/grpc --go-grpc_opt=paths=source_relative ./grpc.proto
