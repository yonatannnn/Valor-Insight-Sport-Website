package usecases

import (
	"time"
	"valorInsight/domain"
	interfaces "valorInsight/domain/interfaces"
	"valorInsight/infrastructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	UserRepository interfaces.UserRepository
}

func NewUserUsecase(ur interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: ur,
	}
}

func (uu *UserUsecase) RegisterUser(user domain.User) (string, domain.Error) {
	_ , err := uu.UserRepository.GetUserByUsername(user.Username)
    if err.Message == "" {
        return "" , err
    }
    if err.Message == "" {
        return "", domain.Error{Message: "Username already exists"}
    }

    hashedPassword, err := infrastructure.CashPassword(user.Password)
    if err.Message != "" {
        return "", err
    }
    user.Password = string(hashedPassword)
    user.ID = primitive.NewObjectID()
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    user.Role = "user"
    user.IsPremium = false
    user.ProfilePicture = ""

    _ , err = uu.UserRepository.RegisterUser(user)
    if err.Message != "" {
        return "", err
    }
    return "User registered", domain.Error{}
}

