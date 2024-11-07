package repositories

import (
	"context"
	"valorInsight/domain"
	interfaces "valorInsight/domain/interfaces"
)

type UserRepository struct {
	dc  interfaces.Collection
	ctx context.Context
}

func NewUserRepository(dc interfaces.Collection, ctx context.Context) *UserRepository {
	return &UserRepository{
		dc:  dc,
		ctx: ctx,
	}
}

func (ur *UserRepository) RegisterUser(user domain.User) (string, domain.Error) {
	var e domain.Error
	_, err := ur.dc.InsertOne(ur.ctx, user)
	if err != nil {
		e.Message = err.Error()
		e.StatusCode = 500
		return "", e
	}
	return "User registered", e
}

func (ur *UserRepository) GetUserByUsername(username string) (domain.User, domain.Error) {
	var user domain.User
	var e domain.Error
	err := ur.dc.FindOne(ur.ctx, domain.User{Username: username}).Decode(&user)
	if err != nil {
		e.Message = err.Error()
		e.StatusCode = 500
		return domain.User{}, e
	}
	return user, e
}
