package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/amajakai14/comment-simple-api/internal/comment"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func convertToComment(req PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   req.Slug,
		Author: req.Author,
		Body:   req.Body,
	}
}

func (h *Handler) PostComment(c *gin.Context) {
	r := c.Request
	w := c.Writer
	var cmtRequest PostCommentRequest

	if err := json.NewDecoder(r.Body).Decode(&cmtRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	validate := validator.New()
	err := validate.Struct(cmtRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	cmt := convertToComment(cmtRequest)

	cmt, err = h.Service.PostComment(r.Context(), cmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(c *gin.Context) {
	id := c.Param("id")

	cmt, err := h.Service.GetComment(c.Request.Context(), id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := json.NewEncoder(c.Writer).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var cmt comment.Comment
	if err := json.NewDecoder(c.Request.Body).Decode(&cmt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmt, err := h.Service.UpdateComment(c.Request.Context(), id, cmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.DeleteComment(c.Request.Context(), id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}
