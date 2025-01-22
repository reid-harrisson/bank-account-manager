package server

import (
	"bank-account-manager/storage"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App     *fiber.App
	Storage *storage.Storage
}

func Create() *Server {
	app := fiber.New()
	storage := storage.Create()

	return &Server{
		App:     app,
		Storage: storage,
	}
}

func (server *Server) Listen(port string) error {
	return server.App.Listen(":" + port)
}
