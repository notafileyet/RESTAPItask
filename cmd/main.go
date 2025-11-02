package main

import (
	"github.com/labstack/echo/v4"

	"APIhendler/internal/config"
	"APIhendler/internal/handlers"
	"APIhendler/internal/tasksRepo"
	"APIhendler/internal/tasksService"
	"APIhendler/internal/userService/repository"
	"APIhendler/internal/userService/service"
	"APIhendler/internal/web/tasks"
	"APIhendler/internal/web/users"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)
	tasksRepo := tasksRepo.NewTaskRepository(db)

	usersService := service.NewUserService(usersRepo, tasksRepo)
	tasksService := tasksService.NewTaskService(tasksRepo, usersRepo)

	tasksHandler := handlers.NewTaskHandlers(tasksService)
	usersHandler := handlers.NewUserHandlers(usersService)

	e := echo.New()

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)

	tasks.RegisterHandlers(e, tasksStrictHandler)
	users.RegisterHandlers(e, usersStrictHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
