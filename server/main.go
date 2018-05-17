package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shuymn/sample-app-with-golang-and-nuxtjs/server/handler"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Initialize handler
	h := &handler.Handler{}

	// routes
	e.GET("/", h.GetIndex)
	e.GET("/users/:id", h.GetUser)

	// start
	e.Logger.Fatal(e.Start(":1323"))
}
