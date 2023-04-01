package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router  *gin.Engine
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = gin.Default()
	h.Router.Use(JSONMiddleware())
	h.Router.Use(TimeoutMiddleware())
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.GET("/alive", func(c *gin.Context) {
		fmt.Fprint(c.Writer, "server is alives")
	})

	h.Router.POST("/api/v1/comment", h.PostComment)
	h.Router.GET("/api/v1/comment/:id", h.GetComment)
	h.Router.PUT("/api/v1/comment/:id", h.UpdateComment)
	h.Router.DELETE("/api/v1/comment/:id", h.DeleteComment)

}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("Server shut down gracefuly")
	return nil
}
