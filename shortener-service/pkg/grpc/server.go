package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"url-shortener/shortener-service/internal/ports"
	"url-shortener/shortener-service/pkg/grpc/urlpb"
)

const address = "localhost:9003"

func NewGRPCAdapter(app ports.AppPort) *Adapter {
	return &Adapter{api: app}
}

func NewURLServiceClient() urlpb.URLServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error did not connect to token service %v", err)
	}

	log.Println("Successfully created url service client")
	return urlpb.NewURLServiceClient(conn)
}

func (grpcA Adapter) GetLongURL(ctx context.Context, req *urlpb.GetLongURLRequest) (*urlpb.GetLongURLResponse, error) {
	longURL, err := grpcA.api.GetURL(req.ShortURL)

	if err != nil {
		return nil, status.Error(codes.Internal, "unexcepted error")
	}

	return &urlpb.GetLongURLResponse{LongURL: longURL}, nil
}

func (grpcA Adapter) GenerateShortURL(ctx context.Context, req *urlpb.NewShortURLRequest) (*urlpb.NewShortURLResponse, error) {
	shortURL, err := grpcA.api.CreateURL(req.LongURL)

	if err != nil {
		return nil, status.Error(codes.Internal, "unexcepted error")
	}

	return &urlpb.NewShortURLResponse{ShortURL: shortURL}, nil
}
