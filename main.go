package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/p886/byo-database/repl"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env: '%s'\n", err.Error())
	}
	storageFilePath := os.Getenv("STORAGE_FILE_PATH")
	log.Printf("Using '%s' as backend", storageFilePath)
	fmt.Println("")

	fmt.Println("Welcome! Enter command prefixed with PUT to store, GET to retrieve.")
	repl.Loop(storageFilePath)
}
