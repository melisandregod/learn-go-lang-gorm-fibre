package main 

import (
	"github.com/gofiber/fiber/v2"
	"jajar-test/handler"
	"jajar-test/service"
)

func SetupRoutes(app *fiber.App, txService *service.TransactionService) {

	txHandler := handler.NewTransactionHandler(txService)
	app.Get("/transactions", txHandler.GetTransactions)
	
}
