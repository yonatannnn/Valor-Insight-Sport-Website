package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name           string             `bson:"name" json:"name"`
	Age            int                `bson:"age" json:"age"`
	Position       string             `bson:"position" json:"position"`
	TeamID         primitive.ObjectID `bson:"team_id" json:"team_id"`
	Nationality    string             `bson:"nationality" json:"nationality"`
	Stats          Stats              `bson:"stats" json:"stats"`
	Bio            string             `bson:"bio" json:"bio"`
	ProfilePicture string             `bson:"profile_picture" json:"profile_picture"`
}

type Stats struct {
	Goals   int `bson:"goals" json:"goals"`
	Assists int `bson:"assists" json:"assists"`
}
