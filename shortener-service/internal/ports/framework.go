package ports

import (
	"context"
	"url-shortener/shortener-service/pkg/grpc/urlpb"
)

type MessageStruct struct {
	Name  string
	Email string
	URL   string
}

type DbPort interface {
	SaveURL(longURL, shortURL string) error
	GetURL(shortURL string) (string, error)
	DeleteURL(id string) error
}

type GRPCPort interface {
	RunServer()
	GenerateShortURL(ctx context.Context, req *urlpb.NewShortURLRequest) (*urlpb.NewShortURLResponse, error)
	GetLongURL(ctx context.Context, req *urlpb.GetLongURLRequest) (*urlpb.GetLongURLResponse, error)
}

type CachePort interface {
	SaveURL(shortURL, longURL string) error
	GetURL(shortURL string) string
}
