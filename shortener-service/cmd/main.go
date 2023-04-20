package main

import (
	"url-shortener/shortener-service/internal/adapters/app"
	"url-shortener/shortener-service/internal/adapters/cache"
	"url-shortener/shortener-service/internal/adapters/core"
	"url-shortener/shortener-service/internal/adapters/db"
	"url-shortener/shortener-service/pkg/grpc"
)

func main() {

	dbAdapter := db.NewAdapter()

	cacheAdapter := redis.NewAdapter()

	coreAdapter := core.NewAdapter()

	appAdapter := app.NewAdapter(coreAdapter, dbAdapter, cacheAdapter)

	grpcA := grpc.NewGRPCAdapter(appAdapter)

	grpcA.RunServer()

}
