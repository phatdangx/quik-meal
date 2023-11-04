package routes

import (
	"user/handler"

	"github.com/labstack/echo/v4"
)

func Private(e *echo.Echo, h *handler.PrivateHandler) {
	privateRoute := e.Group("/v1/private/user")
	privateRoute.GET("/register", h.RegisterAccount)
}
