package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type News struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Author    News_Author        `bson:"author" json:"author"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Tags      []string           `bson:"tags" json:"tags"`
	Image     string             `bson:"image" json:"image"`
}

type News_Author struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name           string             `bson:"name" json:"name"`
	ProfilePicture string             `bson:"profile_picture" json:"profile_picture"`
}
