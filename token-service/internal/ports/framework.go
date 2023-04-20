package ports

import (
	"context"
	"url-shortener/token-service/pkg/grpc/tokenpb"
)

type GRPCPort interface {
	Run()
	ValidateToken(ctx context.Context, req *tokenpb.ValidateTokenRequest) (*tokenpb.ValidateTokenResponse, error)
	GenerateToken(ctx context.Context, req *tokenpb.NewTokenRequest) (*tokenpb.NewTokenResponse, error)
}
