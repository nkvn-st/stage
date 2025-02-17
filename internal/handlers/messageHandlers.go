package handlers

import (
	"context"
	"stage/internal/messageservice"
	"stage/internal/web/messages"
)

type MessageHandler struct {
	Service *messageservice.MessageService
}

func (h *MessageHandler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
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

func (h *MessageHandler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
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

func (h *MessageHandler) PatchMessageById(_ context.Context, request messages.PatchMessageByIdRequestObject) (messages.PatchMessageByIdResponseObject, error) {
	messageRequest := request.Body

	messageToUpdate := messageservice.Message{
		IsDone: *messageRequest.IsDone,
	}
	updatedMessage, err := h.Service.UpdateMessageByID(request.Id, messageToUpdate)

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
	err := h.Service.DeleteMessageByID(request.Id)
	if err != nil {
		return nil, err
	}

	return messages.DeleteMessageById204Response{}, err
}

func NewMessageHandler(service *messageservice.MessageService) *MessageHandler {
	return &MessageHandler{
		Service: service,
	}
}
