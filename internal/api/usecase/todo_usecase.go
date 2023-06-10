package usecase

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/internal/api/repository"
)

type ITodoUsecase interface {
	CreateTodo(req model.Todo) (model.TodoResponse, error)
	ReadTodo(req model.Todo) (model.TodoResponse, error)
	GetAllTodoList(req model.Todo) ([]model.TodoResponse, error)
	UpdateTodo(req model.Todo) (model.TodoResponse, error)
	DeleteTodo(req model.Todo) (model.Message, error)
}

type todoUsecase struct {
	tr repository.ITodoRepository
}

func NewTodoUsecase(tr repository.ITodoRepository) ITodoUsecase {
	return &todoUsecase{tr}
}

func (tu *todoUsecase) CreateTodo(req model.Todo) (model.TodoResponse, error) {

	newTodo := model.Todo{
		Title:  req.Title,
		UserID: req.UserID,
	}

	res, err := tu.tr.CreateTodo(newTodo)
	if err != nil {
		return model.TodoResponse{}, err
	}

	return res, nil

}

func (tu *todoUsecase) ReadTodo(req model.Todo) (model.TodoResponse, error) {

	newTodo := model.Todo{
		ID:     req.ID,
		UserID: req.UserID,
	}

	res, err := tu.tr.ReadTodo(newTodo)
	if err != nil {
		return model.TodoResponse{}, err
	}

	return res, nil

}

func (tu *todoUsecase) GetAllTodoList(req model.Todo) ([]model.TodoResponse, error) {

	newTodo := model.Todo{
		UserID: req.UserID,
	}

	res, err := tu.tr.GetAllTodos(newTodo)
	if err != nil {
		return []model.TodoResponse{}, err
	}

	return res, nil

}

func (tu *todoUsecase) UpdateTodo(req model.Todo) (model.TodoResponse, error) {

	newTodo := model.Todo{
		ID:     req.ID,
		Title:  req.Title,
		UserID: req.UserID,
	}

	res, err := tu.tr.UpdateTodo(newTodo)
	if err != nil {
		return model.TodoResponse{}, err
	}

	return res, nil

}

func (tu *todoUsecase) DeleteTodo(req model.Todo) (model.Message, error) {

	newTodo := model.Todo{
		ID:     req.ID,
		UserID: req.UserID,
	}

	res, err := tu.tr.DeleteTodo(newTodo)
	if err != nil {
		return model.Message{}, err
	}

	return res, nil

}
