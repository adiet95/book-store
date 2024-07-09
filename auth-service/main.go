package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/adiet95/book-store/auth-service/src/config"
	_ "github.com/joho/godotenv/autoload"
)

var initCommand = cobra.Command{
	Short: "Simple backend login & register",
}

func main() {
	args := os.Args[1:]
	log.Println(args)
	if len(args) <= 0 {
		args = []string{"serve"}
	}
	if err := config.Run(args); err != nil {
		log.Fatal(err)
	}
}
