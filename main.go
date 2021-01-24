package main

import (
	"fmt"
	"github.com/WeliHenry/golangblogV1/app"
	"github.com/WeliHenry/golangblogV1/config"
)

func main() {
	fmt.Println("application started successfully")
	MongoURI := config.MongoURI
	fmt.Println("Starting application...")
	app := &app.App{}
	app.Initialize(MongoURI)
	var port = "5000"
	fmt.Println(`Server running @` + port)
	app.Run(":" + port)
}
