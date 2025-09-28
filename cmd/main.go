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

	tasksRepo := tasksRepo.NewTaskRepository(db)
	tasksService := tasksService.NewTaskService(tasksRepo)
	tasksHandler := tasks.NewTaskHandlerAdapter(tasksService)

	usersRepo := repository.NewUserRepository(db)
	usersService := service.NewUserService(usersRepo)
	usersHandler := handlers.NewUserHandlers(usersService)

	e := echo.New()

	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil) // Добавляем строгий хендлер!
	tasks.RegisterHandlers(e, strictTasksHandler)
	
	strictUsersHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, strictUsersHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
