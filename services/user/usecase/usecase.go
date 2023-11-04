package usecase

import (
	"context"
	"user/models"
)

type UserRepository interface {
	CreateNewUser(ctx context.Context, user *models.User) error
}

type usecase struct {
	user UserRepository
}

func NewUserUseCase(userRepo UserRepository) *usecase {
	return &usecase{user: userRepo}
}

func (u *usecase) CreateNewUser(ctx context.Context, user *models.User) error {
	err := u.user.CreateNewUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
