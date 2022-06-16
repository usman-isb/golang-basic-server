package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// routes
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello World")
	})
	e.GET("/contacts", getContacts)
	e.GET("/contacts/:id", getContact)
	e.POST("/contacts", saveContact)
	e.PUT("/contacts/:id", updateContact)
	e.DELETE("/contact/:id", deleteContact)

	// start server
	e.Logger.Fatal(e.Start(":3333"))
}
