package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerificationCode struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email               string             `bson:"email" json:"email"`
	Code                string             `bson:"code" json:"code"`
	ExpiresAt           time.Time          `bson:"expiresAt" json:"expiresAt"`
	VerificationCode_id string             `bson:"verificationCode_id" json:"verificationCode_id"`
}
