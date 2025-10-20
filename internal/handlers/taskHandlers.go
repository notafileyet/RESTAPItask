package handlers

import (
	"context"

	"APIhendler/internal/tasksRepo"
	"APIhendler/internal/tasksService"
	"APIhendler/internal/web/tasks"
)

type TaskHandlers struct {
	TaskService *tasksService.TaskService
}

func NewTaskHandlers(taskService *tasksService.TaskService) *TaskHandlers {
	return &TaskHandlers{TaskService: taskService}
}

func (h *TaskHandlers) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	task := &tasksRepo.Task{
		Title:  taskRequest.Title,
		Status: taskRequest.Status,
		UserID: uint(taskRequest.UserId),
	}

	if err := h.TaskService.CreateTask(task); err != nil {
		return nil, err
	}

	id := int64(task.ID)
	userId := int64(task.UserID)

	response := tasks.PostTasks201JSONResponse{
		Id:        &id,
		Title:     task.Title,
		Status:    task.Status,
		UserId:    userId,
		CreatedAt: &task.CreatedAt,
		UpdatedAt: &task.UpdatedAt,
	}

	return response, nil
}

func (h *TaskHandlers) GetTasks(ctx context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	tasksList, err := h.TaskService.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, t := range tasksList {
		id := int64(t.ID)
		userId := int64(t.UserID)

		task := tasks.Task{
			Id:        &id,
			Title:     t.Title,
			Status:    t.Status,
			UserId:    userId,
			CreatedAt: &t.CreatedAt,
			UpdatedAt: &t.UpdatedAt,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandlers) GetTasksUsersUserId(ctx context.Context, request tasks.GetTasksUsersUserIdRequestObject) (tasks.GetTasksUsersUserIdResponseObject, error) {
	tasksList, err := h.TaskService.GetTasksByUserID(uint(request.UserId))
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksUsersUserId200JSONResponse{}
	for _, t := range tasksList {
		id := int64(t.ID)
		userId := int64(t.UserID)

		task := tasks.Task{
			Id:        &id,
			Title:     t.Title,
			Status:    t.Status,
			UserId:    userId,
			CreatedAt: &t.CreatedAt,
			UpdatedAt: &t.UpdatedAt,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandlers) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body

	task := &tasksRepo.Task{
		ID:     uint(request.Id),
		Title:  taskRequest.Title,
		Status: taskRequest.Status,
		UserID: uint(taskRequest.UserId),
	}

	if err := h.TaskService.UpdateTask(task); err != nil {
		return nil, err
	}

	id := int64(task.ID)
	userId := int64(task.UserID)

	response := tasks.PatchTasksId200JSONResponse{
		Id:        &id,
		Title:     task.Title,
		Status:    task.Status,
		UserId:    userId,
		CreatedAt: &task.CreatedAt,
		UpdatedAt: &task.UpdatedAt,
	}

	return response, nil
}

func (h *TaskHandlers) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	if err := h.TaskService.DeleteTask(uint(request.Id)); err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}
