package main

import (
	"fmt"
	"log"
	"os"
	"valorInsight/infrastructure"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")
	databaseName := "valorInsight"
	userCollection := infrastructure.ConnectMonogodb(databaseName, "users" , mongoURI)
	fmt.Println("Connected to MongoDB collection:", userCollection.Name())

}
