package main

import (
	"github.com/vipul-08/student-api/app"
	"github.com/vipul-08/student-api/db"
)

func main() {
	db.ConnectDb()
	app.StartRoutes()
}
