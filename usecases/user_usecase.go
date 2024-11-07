package usecases

import (
	"valorInsight/domain"
	interfaces "valorInsight/domain/interfaces"
)

type UserUsecase struct {
	UserRepository interfaces.UserRepository
}

func NewUserUsecase(ur interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: ur,
	}
}

func (uu *UserUsecase) RegisterUser(user domain.User) (string, error) {
	uu.UserRepository.RegisterUser(user)
}
