package domain

import (
	"valorInsight/domain"
)

type UserUsecase interface {
	RegisterUser(User domain.User) (string, error)
}
