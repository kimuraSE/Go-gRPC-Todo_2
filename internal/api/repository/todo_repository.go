package repository

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/pkg/todo"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ITodoRepository interface {
	CreateTodo(req model.Todo) (model.TodoResponse, error)
	ReadTodo(req model.Todo) (model.TodoResponse, error)
	GetAllTodos(req model.Todo) ([]model.TodoResponse, error)
	UpdateTodo(req model.Todo) (model.TodoResponse, error)
	DeleteTodo(req model.Todo) (model.Message, error)
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
	md := metadata.New(map[string]string{"authorization": "Bearer test"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := client.Create(ctx, &todo.Todo{
		Title:  req.Title,
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

func (tr *todoRepository) ReadTodo(req model.Todo) (model.TodoResponse, error) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return model.TodoResponse{}, err
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)
	md := metadata.New(map[string]string{"authorization": "Bearer test"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.Read(ctx, &todo.ReadRequest{
		UserId: uint32(req.UserID),
		TodoId: uint32(req.ID),
	})

	if err != nil {
		return model.TodoResponse{}, err
	}

	return model.TodoResponse{
		ID:    uint(res.Id),
		Title: res.Title,
	}, nil

}

func (tr *todoRepository) GetAllTodos(req model.Todo) ([]model.TodoResponse, error) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return []model.TodoResponse{}, err
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)
	md := metadata.New(map[string]string{"authorization": "Bearer test"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.GetAllTodoList(ctx, &todo.TodoListRequest{
		UserId: uint32(req.UserID),
	})

	if err != nil {
		return []model.TodoResponse{}, err
	}

	var todos []model.TodoResponse

	for _, todo := range res.Todos {
		todos = append(todos, model.TodoResponse{
			ID:    uint(todo.Id),
			Title: todo.Title,
		})
	}

	return todos, nil

}

func (tr *todoRepository) UpdateTodo(req model.Todo) (model.TodoResponse, error) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return model.TodoResponse{}, err
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)

	md := metadata.New(map[string]string{"authorization": "Bearer test"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := client.Update(ctx, &todo.UpdateTodoRequest{
		TodoId: uint32(req.ID),
		Title:  req.Title,
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

func (tr *todoRepository) DeleteTodo(req model.Todo) (model.Message, error) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return model.Message{}, err
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)

	md := metadata.New(map[string]string{"authorization": "Bearer test"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := client.Delete(ctx, &todo.DeleteTodoRequest{
		TodoId: uint32(req.ID),
		UserId: uint32(req.UserID),
	})

	if err != nil {
		return model.Message{}, err
	}

	return model.Message{
		Message: res.Message,
	}, nil
}
