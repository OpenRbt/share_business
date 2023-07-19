package main

//go:generate rm -rf ./openapi/restapi ./openapi/client ./openapi/models
//go:generate swagger generate server -t ./openapi/ -f ./openapi/swagger.yaml --strict-responders --strict-additional-properties --principal washBonus/internal/app.Auth --exclude-main
//go:generate swagger generate client -t ./openapi/ -f ./openapi/swagger.yaml --strict-responders --strict-additional-properties --principal washBonus/internal/app.Auth
//go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;
