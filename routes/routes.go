package routes

import (
	_ "bank-account-manager/docs"
	"bank-account-manager/handlers"
	"bank-account-manager/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @Title Bank Account Manager API
// @Version 1.0
// @BasePath /api/v1
// @Description RESTful API endpoints for Bank Account Management
func ConfigRoutes(server *server.Server) {
	server.App.Get("/swagger/*", swagger.HandlerDefault)
	server.App.Get("/", redirectToSwagger)
	apiV1 := server.App.Group("api/v1")

	accountHandler := handlers.CreateAccountHandler(server)

	apiV1.Post("/accounts", accountHandler.Create)
	apiV1.Get("/accounts/:id", accountHandler.ReadOne)
	apiV1.Get("/accounts", accountHandler.ReadAll)

	transactionHandler := handlers.CreateTransactionHandler(server)

	apiV1.Post("/accounts/:id/transactions", transactionHandler.Create)
	apiV1.Get("/accounts/:id/transactions", transactionHandler.ReadByAccount)
	apiV1.Post("/transfer", transactionHandler.Transfer)
}

func redirectToSwagger(context *fiber.Ctx) error {
	return context.Redirect("/swagger/index.html")
}
