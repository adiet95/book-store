package main

import (
	"log"
	"os"

	"github.com/adiet95/book-store/author-service/src/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		args = []string{"serve"}
	}
	if err := config.Run(args); err != nil {
		log.Fatal(err)
	}
}
