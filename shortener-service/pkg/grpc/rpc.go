package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"url-shortener/shortener-service/internal/ports"
	"url-shortener/shortener-service/pkg/grpc/urlpb"
)

type Adapter struct {
	api ports.AppPort
	urlpb.UnimplementedURLServiceServer
}

func (grpcA Adapter) RunServer() {

	listen, err := net.Listen("tcp", ":9003")

	if err != nil {
		log.Fatalf("failed to listening %v", err)
	}

	tokenService := grpcA

	grpcServer := grpc.NewServer()

	urlpb.RegisterURLServiceServer(grpcServer, tokenService)

	log.Println("Url Shortener server is serving at ", listen.Addr())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serving grpc server %v", err)
	}

}
