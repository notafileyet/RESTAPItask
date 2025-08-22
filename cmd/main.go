package main

import (
	"github.com/labstack/echo/v4"

	"APIhendler/internal/config"
	"APIhendler/internal/repository"
	"APIhendler/internal/service"
	"APIhendler/internal/web/tasks"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)

	taskHandler := tasks.NewTaskHandlerAdapter(taskService)

	e := echo.New()

	strictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
