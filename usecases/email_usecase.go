package usecases

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"valorInsight/domain"
	interfaces "valorInsight/domain/interfaces"
	"valorInsight/infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type emailUsecase struct {
	repo         interfaces.VerificationRepository
	emailService infrastructure.EmailService
}

func NewEmailUsecase(repo interfaces.VerificationRepository, emailService infrastructure.EmailService) interfaces.EmailUsecase {
	return &emailUsecase{repo: repo, emailService: emailService}
}

func (u *emailUsecase) SendVerificationCode(ctx context.Context, email string) error {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	expiration := time.Now().Add(10 * time.Minute)
	id := primitive.NewObjectID()
	strId := id.Hex()

	verificationCode := domain.VerificationCode{
		Email:               email,
		Code:                code,
		ExpiresAt:           expiration,
		VerificationCode_id: strId,
	}

	if err := u.repo.SaveCode(ctx, verificationCode); err != nil {
		return err
	}

	return u.emailService.SendVerificationCode(email, code)
}

func (u *emailUsecase) VerifyCode(ctx context.Context, email, code string) error {
	storedCode, err := u.repo.GetCode(ctx, email)
	if err != nil || storedCode == nil {
		return errors.New("verification code not found")
	}

	if storedCode.Code != code || time.Now().After(storedCode.ExpiresAt) {
		return errors.New("invalid or expired verification code")
	}

	return nil
}
