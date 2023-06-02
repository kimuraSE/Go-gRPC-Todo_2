package main

import (
	"Go-REST-Todo/internal/api/controller"
	"Go-REST-Todo/internal/api/repository"
	"Go-REST-Todo/internal/api/usecase"
	"Go-REST-Todo/pkg/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	todoRepository := repository.NewTodoRepository()
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoController := controller.NewTodoController(todoUsecase)

	routes := routes.NewRoutes(userController,todoController)
	routes.Logger.Fatal(routes.Start(":8080"))
}
