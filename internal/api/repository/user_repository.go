package repository

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/pkg/user"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type IUserRepository interface {
	CreateUser(req model.UserRequest) (string, error)
	LoginUser(req model.UserRequest) (string, error)
}

type userRepository struct {
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) CreateUser(req model.UserRequest) (string, error) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)
	md := metadata.New(map[string]string{"authorization": "Bearer test"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.RegisterUser(ctx, &user.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}

	return res.Token, nil
}

func (ur *userRepository) LoginUser(req model.UserRequest) (string, error) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)
	md := metadata.New(map[string]string{"authorization": "Bearer test"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.LoginUser(ctx, &user.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	return res.Token, nil
}
