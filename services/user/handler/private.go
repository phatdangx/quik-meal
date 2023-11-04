package handler

import (
	"context"
	"net/http"
	"user/models"

	"github.com/labstack/echo/v4"
)

type UseCase interface {
	CreateNewUser(ctx context.Context, user *models.User) error
}

type PrivateHandler struct {
	usecase UseCase
}

func NewPrivateHandler(usecase UseCase) *PrivateHandler {
	return &PrivateHandler{
		usecase: usecase,
	}
}

func (h *PrivateHandler) RegisterAccount(c echo.Context) error {
	// bind
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, struct{}{})
}
