package usecase

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/internal/api/repository"
)

type ITodoUsecase interface {
	CreateTodo(req model.Todo) (model.TodoResponse, error)
}

type todoUsecase struct {
	tr repository.ITodoRepository
}

func NewTodoUsecase(tr repository.ITodoRepository) ITodoUsecase {
	return &todoUsecase{tr}
}

func (tu *todoUsecase) CreateTodo(req model.Todo) (model.TodoResponse, error) {

	newTodo := model.Todo{
		Title: req.Title,
		UserID: req.UserID,
	}

	res, err := tu.tr.CreateTodo(newTodo)
	if err != nil {
		return model.TodoResponse{}, err
	}

	return res, nil

}
