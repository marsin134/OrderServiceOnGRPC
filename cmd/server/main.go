package main

import (
	"fmt"
	"log"
	"orderServiceGRPC/internal/config"
)

func main() {
	if err := config.LoadEnvFile(".env"); err != nil {
		log.Printf("Failed to load .env: %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Failed to read .env: %v", err)
	}
	fmt.Println(cfg)
}
