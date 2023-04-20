package main

import (
	"url-shortener/auth-service/internal/adapters/app"
	"url-shortener/auth-service/internal/adapters/core"
	"url-shortener/auth-service/internal/adapters/db"
	"url-shortener/auth-service/pkg/grpc"
	tokenGrpc "url-shortener/token-service/pkg/grpc"
)

func main() {

	dbAdapter := db.NewAdapter()

	coreAdapter := core.NewAdapter()

	appAdapter := app.NewAdapter(coreAdapter, dbAdapter)

	tokenServiceClient := tokenGrpc.NewTokenServiceClient()

	grpcAdapter := grpc.NewGRPCAdapter(appAdapter, tokenServiceClient)

	grpcAdapter.Run()

}
