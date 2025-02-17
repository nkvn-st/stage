package handlers

import (
	"context"
	"stage/internal/userservice"
	"stage/internal/web/users"
)

type UserHandler struct {
	Service *userservice.UserService
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, us := range allUsers {
		user := users.User{
			Id:       &us.ID,
			Email:    &us.Email,
			Password: &us.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userservice.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h *UserHandler) PatchUserById(_ context.Context, request users.PatchUserByIdRequestObject) (users.PatchUserByIdResponseObject, error) {
	userRequest := request.Body

	userToUpdate := userservice.User{
		Email: *userRequest.Email,
	}
	updatedUser, err := h.Service.UpdateUserByID(request.Id, userToUpdate)

	if err != nil {
		return nil, err
	}

	response := users.PatchUserById200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

func (h *UserHandler) DeleteUserById(_ context.Context, request users.DeleteUserByIdRequestObject) (users.DeleteUserByIdResponseObject, error) {
	err := h.Service.DeleteUserByID(request.Id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUserById204Response{}, err
}

func NewUserHandler(service *userservice.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
