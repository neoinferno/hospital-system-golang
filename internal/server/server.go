package server

import (
	"fmt"
	"hospital-system/internal/database"
	"net/http"
)

type Server struct {
	postgres database.Service
	port     int
}

func NewServer() *http.Server {
	postgres := database.NewService()

	NewServer := &Server{
		postgres: postgres,
		port:     8888,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(),
	}
	fmt.Println("Server started on port", NewServer.port)
	return server
}
