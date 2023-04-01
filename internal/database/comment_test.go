//go:build integration
// +build integration

package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/amajakai14/comment-simple-api/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug: "slug",
			Author: "author",
			Body: "body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)


		fmt.Println("test create comment")
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug: "new-slug",
			Author: "means",
			Body: "deletebody",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})
}
