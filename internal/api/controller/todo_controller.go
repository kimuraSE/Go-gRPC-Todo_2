package controller

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/internal/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITodoController interface {
	Create(c echo.Context) error
}

type todoController struct {
	tu usecase.ITodoUsecase
}

func NewTodoController(tu usecase.ITodoUsecase) ITodoController {
	return &todoController{tu}
}

func (tc *todoController) Create(c echo.Context) error {
	newTodo := model.TodoRequest{}
	if err := c.Bind(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// res, err := tc.tu.CreateTodo(newTodo)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	res := model.TodoResponse{
		Title: newTodo.Title,
		UserId: newTodo.ID,
	}

	return c.JSON(http.StatusOK, res)

}
