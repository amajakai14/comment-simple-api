//go:build e2e
// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()
		res, err := client.R().
			SetBody(`{"slug": "/", "author": "Means", "body": "hello from test"}`).
			Post("http://127.0.0.1:8080/api/v1/comment")

		assert.NoError(t, err)
		assert.Equal(t, 201, res.StatusCode())
	})
}
