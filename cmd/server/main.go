package main

import (
	"context"
	"fmt"

	"github.com/amajakai14/comment-simple-api/internal/comment"
	"github.com/amajakai14/comment-simple-api/internal/database"
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
	newComment, err := cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID: "8a30ecaa-2a79-4f46-ad0b-c2f0cfc8bee2",
			Slug: "manual-test",
			Author: "Means",
			Body: "Hello World!",
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newComment)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"8a30ecaa-2a79-4f46-ad0b-c2f0cfc8bee2",
	))

	return nil
}

func main() {
	fmt.Println("Hello, World!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
