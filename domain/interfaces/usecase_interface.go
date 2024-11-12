package domain

import (
	"valorInsight/domain"
)

type UserUsecase interface {
	RegisterUser(User domain.User) (string, domain.Error)
	Login(User domain.User) (string, domain.Error)
}
