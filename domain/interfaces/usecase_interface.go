package domain

import (
	"valorInsight/domain"
)

type UserUsecase interface {
	GetUser() (string, error)
	RegisterUser(User domain.User) (string, error)
}
