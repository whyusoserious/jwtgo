package main

import (
	"context"
	"log"
	"os"

	"testMEDOS/tokens"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")
	port := os.Getenv("PORT")
	secretKey := os.Getenv("SECRET")

	db, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		panic(err)
	}
	err = db.Connect(context.Background())

	if err != nil {
		panic(err)
	}

	token := tokens.NewToken(db, secretKey)

	router := gin.Default()
	handler := tokens.NewHandler(token)
	handler.Register(router)
	router.Run(":" + port)
}
