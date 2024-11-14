package repositories

import (
	"context"
	"valorInsight/domain"
	domainInterfaces "valorInsight/domain/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type verificationRepository struct {
	collection *mongo.Collection
}

func NewVerificationRepository(db *mongo.Collection) domainInterfaces.VerificationRepository {
	return &verificationRepository{collection: db}
}

func (r *verificationRepository) SaveCode(ctx context.Context, code domain.VerificationCode) error {
	opts := options.Replace().SetUpsert(true)
	_, err := r.collection.ReplaceOne(ctx, bson.M{"email": code.Email}, code, opts)
	return err
}

func (r *verificationRepository) GetCode(ctx context.Context, email string) (*domain.VerificationCode, domain.Error) {
	var code domain.VerificationCode
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&code)
	if err != nil {
		return nil, domain.Error{Message: "verification code not found", StatusCode: 404}
	}
	return &code, domain.Error{}
}
