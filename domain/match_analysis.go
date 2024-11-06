package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ImageURL    string `json:"image_url" bson:"image_url"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

type MatchAnalysis struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	MatchID      primitive.ObjectID `json:"match_id" bson:"match_id"`
	AnalysisText string             `json:"analysis_text" bson:"analysis_text"`
	Images       []Image            `json:"images" bson:"images"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	AnalystID    primitive.ObjectID `json:"analyst_id" bson:"analyst_id"`
}
