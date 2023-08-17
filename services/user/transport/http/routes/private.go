package routes

import (
	"user/transport/http"

	"github.com/labstack/echo/v4"
)

func Private(e *echo.Echo, handler *http.UserHandler) {
	privateRoute := e.Group("/v1/private/user")
	privateRoute.GET("/user/:id", handler.GetUserDetail)
}
