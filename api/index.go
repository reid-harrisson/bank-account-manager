package api

import (
	"bank-account-manager/routes"
	s "bank-account-manager/server"
	"net/http"

	"github.com/gofiber/adaptor/v2"
)

var (
	server *s.Server
)

func init() {
	server := s.Create()

	routes.ConfigRoutes(server)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(server.App).ServeHTTP(w, r)
}
