package handlers

import (
	"context"

	"APIhendler/internal/userService/orm"
	"APIhendler/internal/userService/service"
	"APIhendler/internal/web/users"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type UserHandlers struct {
	UserService *service.UserService
}

func NewUserHandlers(userService *service.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

func (h *UserHandlers) GetUsers(ctx context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	usersList, err := h.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, u := range usersList {
		id := int64(u.ID)
		user := users.User{
			Id:        &id,
			Email:     openapi_types.Email(u.Email),
			Password:  u.Password,
			CreatedAt: &u.CreatedAt,
			UpdatedAt: &u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandlers) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	user := &orm.User{
		Email:    string(userRequest.Email),
		Password: userRequest.Password,
	}

	if err := h.UserService.CreateUser(user); err != nil {
		return nil, err
	}

	id := int64(user.ID)
	response := users.PostUsers201JSONResponse{
		Id:        &id,
		Email:     openapi_types.Email(user.Email),
		Password:  user.Password,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return response, nil
}

func (h *UserHandlers) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body

	user := &orm.User{
		ID:       uint(request.Id),
		Email:    string(userRequest.Email),
		Password: userRequest.Password,
	}

	if err := h.UserService.UpdateUser(user); err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:        &request.Id,
		Email:     openapi_types.Email(user.Email),
		Password:  user.Password,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return response, nil
}

func (h *UserHandlers) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	if err := h.UserService.DeleteUser(uint(request.Id)); err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
