package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Task struct {
	ID     int    `json:"id"`
	Text   string `json:"task"`
	Status string `json:"status"`
}

var tasks = make(map[int]Task)
var nextID = 1

func main() {
	e := echo.New()

	e.POST("/task", createTask)
	e.GET("/task", getAllTasks)
	e.PATCH("/task/:id", updateTask)
	e.DELETE("/task/:id", deleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}

func createTask(c echo.Context) error {
	var req struct {
		Text   string `json:"task"`
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "неправильный JSON"})
	}

	task := Task{
		ID:     nextID,
		Text:   req.Text,
		Status: req.Status,
	}
	tasks[nextID] = task
	nextID++

	return c.JSON(http.StatusCreated, task)
}

func getAllTasks(c echo.Context) error {
	var result []Task
	for _, task := range tasks {
		result = append(result, task)
	}
	return c.JSON(http.StatusOK, result)
}

func updateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "неправильный ID"})
	}

	task, exists := tasks[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "задача не найдена"})
	}

	var req struct {
		Text   *string `json:"task"`
		Status *string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "неправильный JSON"})
	}

	if req.Text != nil {
		task.Text = *req.Text
	}
	if req.Status != nil {
		task.Status = *req.Status
	}
	tasks[id] = task

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "неправильный ID"})
	}

	if _, exists := tasks[id]; !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "задача не найдена"})
	}

	delete(tasks, id)
	return c.NoContent(http.StatusNoContent)
}
