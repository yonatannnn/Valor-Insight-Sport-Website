package main

import (
	"context"
	"fmt"
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
	databaseName := "valorInsight"

	userCollection := infrastructure.ConnectMonogodb(databaseName, "users", mongoURI)
	userRepositroy := repositories.NewUserRepository(userCollection, context.TODO())
	UserUsecase := usecases.NewUserUsecase(userRepositroy)

	controller := controllers.NewController(UserUsecase)

	r := router.SetupRouter(controller)
	r.Run(":8080")
	fmt.Println("Connected to MongoDB collection:", userCollection.Name())

}
