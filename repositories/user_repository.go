package repositories

import (
	"context"
	"errors"
	"fmt"
	"valorInsight/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	dc  *mongo.Collection
	ctx context.Context
}

func NewUserRepository(dc *mongo.Collection, ctx context.Context) *UserRepository {
	return &UserRepository{
		dc:  dc,
		ctx: ctx,
	}
}

func (ur *UserRepository) RegisterUser(user domain.User) (string, domain.Error) {
	var e domain.Error
	_, err := ur.GetUserByUsername(user.Username)
	fmt.Println(err.Message)
	if err.Message != "User not found" {
		e.Message = "Username already existss"
		e.StatusCode = 400
		return "", e
	}

	_, error := ur.dc.InsertOne(ur.ctx, user)
	if error != nil {
		e.Message = error.Error()
		e.StatusCode = 500
		return "", e
	}

	return "User registered", e
}

func (ur *UserRepository) GetUserByUsername(username string) (domain.User, domain.Error) {
	var user domain.User
	var e domain.Error

	err := ur.dc.FindOne(ur.ctx, bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			e.Message = "User not found"
			e.StatusCode = 404
		} else {
			e.Message = err.Error()
			e.StatusCode = 500
		}
		return domain.User{}, e
	}
	return user, e
}

func (ur *UserRepository) GetUserByID(id string) (domain.User, domain.Error) {
	var user domain.User
	var e domain.Error
	objectID, error := primitive.ObjectIDFromHex(id)
	if error != nil {
		e.Message = "Invalid user ID"
		e.StatusCode = 400
		return domain.User{}, e
	}
	err := ur.dc.FindOne(ur.ctx, domain.User{ID: objectID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			e.Message = "User not found"
			e.StatusCode = 404
		} else {
			e.Message = err.Error()
			e.StatusCode = 500
		}
		return domain.User{}, e
	}
	return user, e
}

func (ur *UserRepository) UpdateUser(user domain.User) domain.Error {
	var e domain.Error
	_, error := ur.dc.ReplaceOne(ur.ctx, bson.D{{"user_id", user.UserId}}, user)
	if error != nil {
		e.Message = error.Error()
		e.StatusCode = 500
		return e
	}
	return e
}
