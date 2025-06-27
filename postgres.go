package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=mypassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	db.AutoMigrate(&Task{})

	e := echo.New()

	e.POST("/tasks", createTask)
	e.GET("/tasks", getTasks)
	e.GET("/tasks/:id", getTaskByID)
	e.PATCH("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}

func createTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "не удалось прочитать данные"})
	}
	db.Create(&task)
	return c.JSON(http.StatusCreated, task)
}

func getTasks(c echo.Context) error {
	var tasks []Task
	db.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

func getTaskByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var task Task
	result := db.First(&task, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "задача не найдена"})
	}
	return c.JSON(http.StatusOK, task)
}

func updateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "задача не найдена"})
	}

	var updateData Task
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ошибка ввода"})
	}

	task.Title = updateData.Title
	task.Status = updateData.Status
	db.Save(&task)

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := db.Delete(&Task{}, id)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "задача не найдена"})
	}
	return c.NoContent(http.StatusNoContent)
}
