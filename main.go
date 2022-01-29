package main

import (
	"fmt"
	gDB "keyserver/database"
	"log"

	restAPI "keyserver/cmd/keyserver-server"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello")
	log.Println("Welcome")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env configuration file")
	}
	if gDB.Connect() {
		restAPI.Begin()
	}
	return
}
