package sharebusiness

//go:generate rm -rf ./internal/api/restapi
//go:generate C:\Users\Kronu\Downloads\swagger_windows_amd64.exe generate server -t ./transport/rest/restapi -f ./swagger.yaml --exclude-main
//go:generate C:\Users\Kronu\Downloads\swagger_windows_amd64.exe generate client -t ./transport/rest/restapi -f ./swagger.yaml

//go:generate protoc --go_out=./transport/grpc --go_opt=paths=source_relative --go-grpc_out=./transport/grpc --go-grpc_opt=paths=source_relative ./transport/grpc.proto
