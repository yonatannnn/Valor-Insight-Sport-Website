package main

import (
	"context"
	"log"
	"os"
	"valorInsight/controllers"
	"valorInsight/infrastructure"
	"valorInsight/repositories"
	"valorInsight/router"
	"valorInsight/usecases"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpEmail := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")

	userCollection := infrastructure.ConnectMonogodb(databaseName, "users", mongoURI)
	verificationCodeCollection := infrastructure.ConnectMonogodb(databaseName, "verification_codes", mongoURI)

	userRepositroy := repositories.NewUserRepository(userCollection, context.TODO())
	emailRepositroy := repositories.NewVerificationRepository(verificationCodeCollection)

	emailService := infrastructure.NewEmailService(smtpEmail, smtpPassword, smtpHost, smtpPort)

	userUsecase := usecases.NewUserUsecase(userRepositroy)
	emailUsecase := usecases.NewEmailUsecase(emailRepositroy, emailService)

	controller := controllers.NewController(userUsecase, emailUsecase)

	r := router.SetupRouter(controller)
	r.Run(":8080")

}
