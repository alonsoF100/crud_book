package handlers

import (
	"crud_book/internal/services"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	bookHandler *bookHandler
	userHandler *userHandler
}

func NewRouter(bookService services.BookService, userService services.UserService) *gin.Engine {
	h := &handlers{
		bookHandler: newBookHandler(bookService),
		userHandler: newUserHandler(userService),
	}

	return h.setupRouter()
}

func (h *handlers) setupRouter() *gin.Engine {
	r := gin.Default()

	// End-points for book
	r.POST("/books", h.bookHandler.createBook)
	r.GET("/books/:id", h.bookHandler.getBook)
	r.PUT("/books/:id/status", h.bookHandler.updateBookStatus)
	r.PUT("/books/:id/rating", h.bookHandler.updateBookRating)
	r.DELETE("/books/:id", h.bookHandler.deleteBook)
	r.GET("/users/:id/books", h.bookHandler.getUserBooks)

	// End-points for user
	r.POST("/users", h.userHandler.createUser)
	r.GET("/users", h.userHandler.getAllUsers)
	r.GET("/users/:id", h.userHandler.getUser)
	r.DELETE("/users/:id", h.userHandler.deleteUser)

	return r
}
