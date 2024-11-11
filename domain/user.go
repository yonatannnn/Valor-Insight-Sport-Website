package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Username       string               `bson:"username" json:"username"`
	Email          string               `bson:"email" json:"email"`
	Password       string               `bson:"password" json:"password"`
	CreatedAt      time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time            `bson:"updated_at" json:"updated_at"`
	FavoriteTeams  []primitive.ObjectID `bson:"favorite_teams" json:"favorite_teams"`
	Role           string               `bson:"role" json:"role"`
	ProfilePicture string               `bson:"profile_picture" json:"profile_picture"`
	IsPremium      bool                 `bson:"is_premiun" json:"is_premium"`
	RefreshToken   string               `bson:"refresh_token" json:"refresh_token"`
}
