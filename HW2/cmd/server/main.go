package main

import (
	"awesomeProject/accounts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := setupServer()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func setupServer() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	setupRoutes(e)

	return e
}

func setupRoutes(e *echo.Echo) {
	accountsHandler := accounts.New()

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.DELETE("/account/delete", accountsHandler.DeleteAccount)
	e.PATCH("/account/patch", accountsHandler.PatchAccount)
	e.PUT("/account/change", accountsHandler.ChangeAccount)
}
