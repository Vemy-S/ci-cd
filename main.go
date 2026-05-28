package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	client, err := ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Disconnect failed: %v", err)
		}
	}()

	log.Println("Application is running smoothly...")
}
