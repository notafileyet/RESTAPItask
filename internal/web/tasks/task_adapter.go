package tasks

import (
	"context"

	"APIhendler/internal/tasksRepo"
	"APIhendler/internal/tasksService"
)

type TaskHandlerAdapter struct {
	service *tasksService.TaskService
}

func NewTaskHandlerAdapter(svc *tasksService.TaskService) *TaskHandlerAdapter {
	return &TaskHandlerAdapter{service: svc}
}

func (a *TaskHandlerAdapter) GetTasks(ctx context.Context, _ GetTasksRequestObject) (GetTasksResponseObject, error) {
	tasksList, err := a.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := GetTasks200JSONResponse{}
	for _, t := range tasksList {
		id := int64(t.ID)
		task := Task{
			Id:     &id,
			Title:  t.Title,
			Status: t.Status,
		}
		response = append(response, task)
	}

	return response, nil
}

func (a *TaskHandlerAdapter) PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error) {
	taskRequest := request.Body

	task := &tasksRepo.Task{
		Title:  taskRequest.Title,
		Status: taskRequest.Status,
	}

	if err := a.service.CreateTask(task); err != nil {
		return nil, err
	}

	id := int64(task.ID)
	response := PostTasks201JSONResponse{
		Id:     &id,
		Title:  task.Title,
		Status: task.Status,
	}

	return response, nil
}

func (a *TaskHandlerAdapter) PatchTasksId(ctx context.Context, request PatchTasksIdRequestObject) (PatchTasksIdResponseObject, error) {
	taskRequest := request.Body

	task := &tasksRepo.Task{
		ID:     uint(request.Id),
		Title:  taskRequest.Title,
		Status: taskRequest.Status,
	}

	if err := a.service.UpdateTask(task); err != nil {
		return nil, err
	}

	response := PatchTasksId200JSONResponse{
		Id:     &request.Id,
		Title:  task.Title,
		Status: task.Status,
	}

	return response, nil
}

func (a *TaskHandlerAdapter) DeleteTasksId(ctx context.Context, request DeleteTasksIdRequestObject) (DeleteTasksIdResponseObject, error) {
	if err := a.service.DeleteTask(uint(request.Id)); err != nil {
		return nil, err
	}

	return DeleteTasksId204Response{}, nil
}
