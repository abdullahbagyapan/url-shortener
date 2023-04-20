package main

import (
	"url-shortener/token-service/internal/adapters/app"
	"url-shortener/token-service/internal/adapters/core"
	"url-shortener/token-service/pkg/grpc"
)

func main() {

	coreAdapter := core.NewCoreAdapter()

	appAdapter := app.NewAppAdapter(coreAdapter)

	grpcAdapter := grpc.NewGRPCAdapter(appAdapter)

	grpcAdapter.Run()

}
