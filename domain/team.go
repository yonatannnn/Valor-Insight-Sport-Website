package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string               `bson:"name" json:"name"`
	Country   string               `bson:"country" json:"country"`
	League    string               `bson:"league" json:"league"`
	Founded   time.Time            `bson:"founded" json:"founded"`
	Stadium   string               `bson:"stadium" json:"stadium"`
	Players   []primitive.ObjectID `bson:"players" json:"players"`
	Logo      string               `bson:"logo" json:"logo"`
	FansCount int                  `bson:"fans_count" json:"fans_count"`
}
