package main

import (
	"Todo_App/internal/config"
	"log"
)

func main() {
	_, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load Configuration:", err)
	}
}
