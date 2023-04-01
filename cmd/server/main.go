package main

import (
	"fmt"

	"github.com/amajakai14/comment-simple-api/internal/comment"
	"github.com/amajakai14/comment-simple-api/internal/database"
	transportHttp "github.com/amajakai14/comment-simple-api/internal/transport/http"
)

// Run - is going to be responsible for 
// the instantiation and startup of our
// go application

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
	if err := httpHandler.Serve(); err != nil {
		fmt.Println("Failed to serve http")
		return err
	}

	fmt.Println("shutting down our application")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
