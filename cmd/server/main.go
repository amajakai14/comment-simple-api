package main

import (
	"fmt"

	"github.com/amajakai14/comment-simple-api/internal/comment"
	"github.com/amajakai14/comment-simple-api/internal/database"
	transportHttp "github.com/amajakai14/comment-simple-api/internal/transport/http"
)

func Run() error {
	fmt.Println("starting up our application")
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}
	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)

	return nil
}

func main() {
	fmt.Println("Hello, World!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
