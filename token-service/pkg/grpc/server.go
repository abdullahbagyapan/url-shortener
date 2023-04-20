package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"url-shortener/token-service/internal/ports"
	tokenpb "url-shortener/token-service/pkg/grpc/tokenpb"
)

type Adapter struct {
	api ports.AppPort
	tokenpb.UnimplementedTokenServiceServer
}

func (grpcA Adapter) Run() {

	listen, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listening %v", err)
	}

	tokenService := grpcA

	grpcServer := grpc.NewServer()

	tokenpb.RegisterTokenServiceServer(grpcServer, tokenService)

	log.Println("Token server is serving at : ", listen.Addr())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serving grpc server %v", err)
	}

}
