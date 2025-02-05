package config

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

var ESClient *elasticsearch.Client

func InitElastic() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or cannot be loaded.")
	}

	esURL := os.Getenv("ELASTIC_URL")
	esAPIKey := os.Getenv("ELASTIC_API_KEY") 

	// ตรวจสอบว่า ELASTIC_URL และ API Key มีค่าหรือไม่
	if esURL == "" || esAPIKey == "" {
		log.Fatal("ELASTIC_URL or ELASTIC_API_KEY is not set in .env file")
	}

	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
		APIKey:    esAPIKey, 
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error connecting to Elasticsearch: %s", err)
	}

	ESClient = client
	log.Println("Connected to Elasticsearch using API Key")
}
