package main

import (
	"github.com/labstack/echo/v4"

	"APIhendler/internal/config"
	"APIhendler/internal/handlers"
	"APIhendler/internal/tasksRepo"
	"APIhendler/internal/tasksService"
	"APIhendler/internal/web/tasks"
	"APIhendler/internal/web/users"

	userServiceRepo "APIhendler/internal/userService/repository"
	userServiceService "APIhendler/internal/userService/service"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	tasksRepo := tasksRepo.NewTaskRepository(db)
	tasksService := tasksService.NewTaskService(tasksRepo)
	tasksHandler := handlers.NewTaskHandlers(tasksService)

	usersRepo := userServiceRepo.NewUserRepository(db, tasksRepo)
	usersService := userServiceService.NewUserService(usersRepo, tasksRepo)
	usersHandler := handlers.NewUserHandlers(usersService)

	e := echo.New()

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)

	tasks.RegisterHandlers(e, tasksStrictHandler)
	users.RegisterHandlers(e, usersStrictHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
