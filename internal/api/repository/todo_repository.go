package repository

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/pkg/todo"
	"context"

	"google.golang.org/grpc"
)

type ITodoRepository interface {
	CreateTodo(req model.Todo) (model.TodoResponse, error)
}

type todoRepository struct {
}

func NewTodoRepository() ITodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) CreateTodo(req model.Todo) (model.TodoResponse, error) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return model.TodoResponse{}, err
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)

	res, err := client.Create(context.Background(), &todo.Todo{
		Title: req.Title,
		UserId: uint32(req.UserID),
	})

	if err != nil {
		return model.TodoResponse{}, err
	}

	return model.TodoResponse{
		ID:    uint(res.Id),
		Title: res.Title,
	}, nil

}
