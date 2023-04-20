package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"url-shortener/auth-service/internal/ports"
	"url-shortener/auth-service/pkg/grpc/pb"
	"url-shortener/shortener-service/pkg/grpc/urlpb"
	"url-shortener/token-service/pkg/grpc/tokenpb"
)

type Adapter struct {
	api ports.AppPort
	pb.UnimplementedAuthServiceServer
	urlService   urlpb.URLServiceClient
	tokenService tokenpb.TokenServiceClient
}

func (grpcA Adapter) Run() {

	listen, err := net.Listen("tcp", ":9001")

	if err != nil {
		log.Fatalf("failed to listening %v", err)
	}

	authService := grpcA

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, authService)

	log.Println("Auth server is serving at : ", listen.Addr())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serving grpc server %v", err)
	}

}
