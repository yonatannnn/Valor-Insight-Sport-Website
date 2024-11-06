package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Match struct {
	TeamAID primitive.ObjectID `json:"team_a_id" bson:"team_a_id"`
	TeamBID primitive.ObjectID `json:"team_b_id" bson:"team_b_id"`
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Date    time.Time          `json:"date" bson:"date"`
	Score   Score              `json:"score" bson:"score"`
	Stadium string             `json:"stadium" bson:"stadium"`
	League  primitive.ObjectID `json:"league" bson:"league"`
}
type Score struct {
	TeamA int `json:"team_a" bson:"team_a"`
	TeamB int `json:"team_b" bson:"team_b"`
}
