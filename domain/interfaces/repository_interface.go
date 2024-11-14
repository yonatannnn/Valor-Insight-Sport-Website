package domain

import (
	"context"
	"valorInsight/domain"
)

type UserRepository interface {
	RegisterUser(user domain.User) (string, domain.Error)
	// Login(username, password string) (domain.User, domain.Error)
	// PromoteUser(userID int) domain.Error
	GetUserByID(id string) (domain.User, domain.Error)
	GetUserByUsername(username string) (domain.User, domain.Error)
	UpdateUser(user domain.User) domain.Error
}

type VerificationRepository interface {
	SaveCode(ctx context.Context, code domain.VerificationCode) error
	GetCode(ctx context.Context, email string) (*domain.VerificationCode, error)
}
