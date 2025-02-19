package main

import (
	"log"
	"stage/internal/database"
	"stage/internal/handlers"
	"stage/internal/messageservice"
	"stage/internal/userservice"
	"stage/internal/web/messages"
	"stage/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&userservice.User{}, &messageservice.Message{})
	if err != nil {
		log.Fatalf("Migration err: %v", err)
	}

	messageRepo := messageservice.NewMessageRepository(database.DB)
	messageService := messageservice.NewService(messageRepo)

	userRepo := userservice.NewUserRepository(database.DB)
	userService := userservice.NewService(userRepo)

	messageHandler := handlers.NewMessageHandler(messageService, userService)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictMessageHandler := messages.NewStrictHandler(messageHandler, nil)
	messages.RegisterHandlers(e, strictMessageHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
