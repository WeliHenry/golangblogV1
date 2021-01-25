package main

import (
	"fmt"
	"github.com/WeliHenry/golangblogV1/app"
)

func main() {
	fmt.Println("application started successfully")
	MongoURI := "mongodb://localhost:27017"
	fmt.Println("Starting application...")
	app := &app.App{}
	app.Initialize(MongoURI)
	var port = "8000"
	fmt.Println(`Server running @` + port)
	app.Run(":" + port)
}
