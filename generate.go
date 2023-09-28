package main

//go:generate rm -rf ./openapi/bonus/* ./openapi/admin/*
//go:generate swagger generate server -t ./openapi/bonus -f ./openapi/bonus.swagger.yaml --strict-responders --strict-additional-properties --principal washbonus/internal/app.Auth --exclude-main
//go:generate swagger generate client -t ./openapi/bonus -f ./openapi/bonus.swagger.yaml --strict-responders --strict-additional-properties --principal washbonus/internal/app.Auth
//go:generate swagger generate server -t ./openapi/admin -f ./openapi/admin.swagger.yaml --strict-responders --strict-additional-properties --principal washbonus/internal/app.AdminAuth --exclude-main
//go:generate swagger generate client -t ./openapi/admin -f ./openapi/admin.swagger.yaml --strict-responders --strict-additional-properties --principal washbonus/internal/app.AdminAuth
//go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;
