package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Prediction struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	MatchID                primitive.ObjectID `bson:"match_id,omitempty" json:"match_id"`
	PredictedWinner        primitive.ObjectID `bson:"predicted_winner,omitempty" json:"predicted_winner"`
	TeamAWinningPercentage float64            `bson:"team_a_winning_percentage,omitempty" json:"team_a_winning_percentage"`
	TeamBWinningPercentage float64            `bson:"team_b_winning_percentage,omitempty" json:"team_b_winning_percentage"`
	DrawPercentage         float64            `bson:"draw_percentage,omitempty" json:"draw_percentage"`
	CreatedAt              time.Time          `bson:"created_at,omitempty" json:"created_at"`
	EndAt                  time.Time          `bson:"end_at,omitempty" json:"end_at"`
	Status                 string             `bson:"status,omitempty" json:"status"`
}
