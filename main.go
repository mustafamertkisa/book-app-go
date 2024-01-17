package main

import (
	"book-app-go/common/app"
	"book-app-go/common/postgresql"
	"book-app-go/controller"
	"book-app-go/database/repository"
	"book-app-go/service"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()

	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	bookRepository := repository.NewBookRepository(dbPool)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	bookController.RegisterRoutes(e)

	e.Start("localhost:8080")
}
