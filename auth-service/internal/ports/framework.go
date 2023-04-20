package ports

import (
	"context"
	"url-shortener/auth-service/pkg/grpc/pb"
)

type User struct {
	Id       string
	Name     string
	Username string
	Password []byte
	Email    string
}

type DbPort interface {
	FindByUsername(username string) (*User, error)
	SaveUser(user *User) error
	CloseDbConnection()
}

type GRPC interface {
	Run()
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.RegisterLoginResponse, error)
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterLoginResponse, error)
}
