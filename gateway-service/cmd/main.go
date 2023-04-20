package main

import (
	authGrpc "url-shortener/auth-service/pkg/grpc"
	"url-shortener/gateway-service/internal/adapters/app"
	"url-shortener/gateway-service/internal/adapters/redis"
	urlGrpc "url-shortener/shortener-service/pkg/grpc"
	tokenGrpc "url-shortener/token-service/pkg/grpc"
)

func main() {

	tokenServiceClient := tokenGrpc.NewTokenServiceClient()

	authServiceClient := authGrpc.NewAuthServiceClient()

	urlServiceClient := urlGrpc.NewURLServiceClient()

	redisClient := redis.NewAdapter()

	appAdapter := app.NewAdapter(authServiceClient, tokenServiceClient, redisClient, urlServiceClient)

	appAdapter.Run()
}
