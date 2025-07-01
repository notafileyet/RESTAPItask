package main

import (
	"github.com/labstack/echo/v4"

	"APIhendler/internal/config"
	"APIhendler/internal/handlers"
	"APIhendler/internal/repository"
	"APIhendler/internal/service"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	e := echo.New()

	e.POST("/tasks", taskHandler.CreateTask)
	e.GET("/tasks", taskHandler.GetAllTasks)
	e.PATCH("/tasks/:id", taskHandler.UpdateTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
