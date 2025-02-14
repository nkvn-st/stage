package main

import (
	"log"
	"stage/internal/database"
	"stage/internal/handlers"
	"stage/internal/messageservice"
	"stage/internal/web/messages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messageservice.Message{})

	repo := messageservice.NewMessageRepository(database.DB)
	service := messageservice.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := messages.NewStrictHandler(handler, nil)
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
