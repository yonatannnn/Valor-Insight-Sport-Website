package domain

import (
	"context"
	"valorInsight/domain"
)

type UserUsecase interface {
	RegisterUser(User domain.User) (string, domain.Error)
	Login(User domain.User) (string, domain.Error)
}

type EmailUsecase interface {
	SendVerificationCode(ctx context.Context, email string) error
	VerifyCode(ctx context.Context, email, code string) error
}
