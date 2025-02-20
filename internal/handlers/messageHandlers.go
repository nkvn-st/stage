package handlers

import (
	"context"
	"stage/internal/messageservice"
	"stage/internal/userservice"
	"stage/internal/web/messages"
)

type MessageHandler struct {
	MessageService *messageservice.MessageService
	UserService    *userservice.UserService
}

func (h *MessageHandler) GetMessagesByUserId(_ context.Context, request messages.GetMessagesByUserIdRequestObject) (messages.GetMessagesByUserIdResponseObject, error) {
	allMessages, err := h.UserService.GetMessagesForUser(request.UserId)
	if err != nil {
		return nil, err
	}

	response := messages.GetMessagesByUserId200JSONResponse{}

	for _, mes := range allMessages {
		message := messages.Message{
			Id:     &mes.ID,
			Task:   &mes.Task,
			IsDone: &mes.IsDone,
			UserId: &mes.UserID,
		}
		response = append(response, message)
	}

	return response, nil
}

func (h *MessageHandler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body

	messageToCreate := messageservice.Message{
		Task:   *messageRequest.Task,
		IsDone: *messageRequest.IsDone,
		UserID: *messageRequest.UserId,
	}
	createdMessage, err := h.MessageService.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:     &createdMessage.ID,
		Task:   &createdMessage.Task,
		IsDone: &createdMessage.IsDone,
		UserId: &createdMessage.UserID,
	}

	return response, nil
}

func (h *MessageHandler) PatchMessageById(_ context.Context, request messages.PatchMessageByIdRequestObject) (messages.PatchMessageByIdResponseObject, error) {
	messageRequest := request.Body

	messageToUpdate := messageservice.Message{
		IsDone: *messageRequest.IsDone,
	}
	updatedMessage, err := h.MessageService.UpdateMessageByID(request.Id, messageToUpdate)

	if err != nil {
		return nil, err
	}

	response := messages.PatchMessageById200JSONResponse{
		Id:     &updatedMessage.ID,
		Task:   &updatedMessage.Task,
		IsDone: &updatedMessage.IsDone,
	}

	return response, nil
}

func (h *MessageHandler) DeleteMessageById(_ context.Context, request messages.DeleteMessageByIdRequestObject) (messages.DeleteMessageByIdResponseObject, error) {
	err := h.MessageService.DeleteMessageByID(request.Id)
	if err != nil {
		return nil, err
	}

	return messages.DeleteMessageById204Response{}, err
}

func NewMessageHandler(messageService *messageservice.MessageService, userService *userservice.UserService) *MessageHandler {
	return &MessageHandler{
		MessageService: messageService,
		UserService:    userService,
	}
}
