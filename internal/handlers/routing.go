package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	// End-points for book
	r.POST("/books", h.CreateBook)
	r.GET("/books/:id", h.GetBook)
	r.PUT("/books/:id/status", h.UpdateBookStatus)
	r.PUT("/books/:id/rating", h.UpdateBookRating)
	r.DELETE("/books/:id", h.DeleteBook)
	r.GET("/users/:id/books", h.GetUserBooks)

	// End-points for user
	r.POST("/users", h.CreateUser)
	r.GET("/users", h.GetAllUsers)
	r.GET("/users/:id", h.GetUser)
	r.DELETE("/users/:id", h.DeleteUser)

	return r
}
