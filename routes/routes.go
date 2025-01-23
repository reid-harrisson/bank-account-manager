package routes

import (
	_ "bank-account-manager/docs"
	"bank-account-manager/handlers"
	"bank-account-manager/server"

	"github.com/gofiber/swagger"
)

// @Title Fiber Swagger Example API
// @Version 1.0
// @Description This is a sample server server.
// @License.name Apache 2.0
// @Host localhost:8080
// @BasePath /api/v1/
func ConfigRoutes(server *server.Server) {
	server.App.Get("/swagger/*", swagger.HandlerDefault)

	apiV1 := server.App.Group("api/v1")

	accountHandler := handlers.CreateAccountHandler(server)

	apiV1.Post("/accounts", accountHandler.Create)
	apiV1.Get("/accounts/:id", accountHandler.ReadOne)
	apiV1.Get("/accounts", accountHandler.ReadAll)
}
