package main

import (
	"context"
	"fmt"
	"log"

	"github.com/luis-mtzz/habits/internal/adapters/api/discogs"
	"github.com/luis-mtzz/habits/pkg/config"
)

func main() {
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	client := discogs.NewClient(config.Discogs.APIKey)

	albums, err := client.GetUserCollection(context.Background(), "it0to")
	if err != nil {
		log.Fatal("Error getting user collection:", err)
	}

	for _, album := range albums {
		fmt.Printf("Title: %s Artist: %s Genre: %s\n", album.Title, album.Artist, album.Genre)
	}
}
