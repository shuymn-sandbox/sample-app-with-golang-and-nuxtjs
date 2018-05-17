package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetIndex handler
func (h *Handler) GetIndex(c echo.Context) error {
	r := &Response{Message: "Hello, world!"}
	return c.JSON(http.StatusOK, r)
}
