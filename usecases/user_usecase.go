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
	hashedPassword, err := infrastructure.CashPassword(user.Password)
	if err.Message != "" {
		return "", err
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Role = "user"
	user.IsPremium = false
	user.ProfilePicture = ""
	user.UserId = user.ID.Hex()

	token, refreshToken, tokenError := infrastructure.GenerateJWT(user)
	if tokenError.Message != "" {
		return "", domain.Error{Message: err.Message, StatusCode: 500}
	}

	user.RefreshToken = refreshToken

	_, err = uu.UserRepository.RegisterUser(user)
	if err.Message != "" {
		return "", err
	}

	return token, domain.Error{}
}
