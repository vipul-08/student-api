package main

import (
	"github.com/joho/godotenv"
	"github.com/vipul-08/student-api/app"
	"github.com/vipul-08/student-api/db"
	"os"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "prod" {
		err := godotenv.Load()
		if err != nil {
			panic("Unable to load .env file")
		}
	}
	db.ConnectDb()
	app.StartRoutes()
}
