package handlers

import (
	"context"
	"stage/internal/messageservice"
	"stage/internal/web/messages"
)

type Handler struct {
	Service *messageservice.MessageService
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	for _, mes := range allMessages {
		message := messages.Message{
			Id:     &mes.ID,
			Task:   &mes.Task,
			IsDone: &mes.IsDone,
		}
		response = append(response, message)
	}

	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body

	messageToCreate := messageservice.Message{
		Task:   *messageRequest.Task,
		IsDone: *messageRequest.IsDone,
	}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:     &createdMessage.ID,
		Task:   &createdMessage.Task,
		IsDone: &createdMessage.IsDone,
	}

	return response, nil
}

func (h *Handler) PatchMessagesId(_ context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	messageRequest := request.Body

	messageToUpdate := messageservice.Message{
		IsDone: *messageRequest.IsDone,
	}
	updatedMessage, err := h.Service.UpdateMessageByID(request.Id, messageToUpdate)

	if err != nil {
		return nil, err
	}

	response := messages.PatchMessagesId200JSONResponse{
		Id:     &updatedMessage.ID,
		Task:   &updatedMessage.Task,
		IsDone: &updatedMessage.IsDone,
	}

	return response, nil
}

func (h *Handler) DeleteMessagesId(_ context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	err := h.Service.DeleteMessageByID(request.Id)
	if err != nil {
		return nil, err
	}

	return messages.DeleteMessagesId204Response{}, err
}

func NewHandler(service *messageservice.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}
