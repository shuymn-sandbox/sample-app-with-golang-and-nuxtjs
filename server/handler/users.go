package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/shuymn/sample-app-with-golang-and-nuxtjs/server/model"
)

// GetUser receiver
func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid id"}
	}
	u := &model.User{ID: id}
	return c.JSON(http.StatusOK, u)
}
