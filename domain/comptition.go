package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Competition struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Country string             `bson:"country" json:"country"`
	Teams   []string           `bson:"teams" json:"teams"`
}
