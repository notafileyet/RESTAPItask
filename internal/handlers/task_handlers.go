package handler

import (
	"net/http"
	"strconv"

	"APIhendler/internal/repository"
	"APIhendler/internal/service"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	Service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var task repository.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный формат"})
	}
	if err := h.Service.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось создать"})
	}
	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при получении"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный ID"})
	}
	task, err := h.Service.GetTaskByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
	}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ошибка привязки"})
	}
	if err := h.Service.UpdateTask(task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка обновления"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный ID"})
	}
	if err := h.Service.DeleteTask(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка удаления"})
	}
	return c.NoContent(http.StatusNoContent)
}
