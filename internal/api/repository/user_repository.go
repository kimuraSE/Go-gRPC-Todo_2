package repository

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/pkg/user"
	"context"

	"google.golang.org/grpc"
)


type IUserRepository interface {
	CreateUser(req model.UserRequest) (string,error)
}

type userRepository struct {
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) CreateUser(req model.UserRequest) (string,error) {
	
	conn,err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	res,err := client.RegisterUser(context.Background(), &user.RegisterRequest{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		})
	if err != nil {
		return "", err
	}

	return res.Token, nil
}