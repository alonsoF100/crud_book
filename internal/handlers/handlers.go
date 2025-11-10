package handlers

import (
	"crud_book/internal/dto"
	"crud_book/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	bookService *services.BookService
	userService *services.UserService
}

func New(bookService *services.BookService, userService *services.UserService) *Handler {
	return &Handler{
		bookService: bookService,
		userService: userService,
	}
}

// Book handlers

func (h *Handler) CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	book, err := h.bookService.CreateBook(req.UserID, req.Name, req.Description)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(201, response)
}

func (h *Handler) GetUserBooks(c *gin.Context) {
	userID := c.Param("id")

	books, err := h.bookService.GetUserBooks(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.BooksToResponse(books)
	c.JSON(200, response)
}

func (h *Handler) UpdateBookStatus(c *gin.Context) {
	var req dto.UpdateBookStatusRequest

	bookID := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	book, err := h.bookService.UpdateBookStatus(bookID, req.Status)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}

func (h *Handler) UpdateBookRating(c *gin.Context) {
	var req dto.UpdateBookRatingRequest

	bookID := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	book, err := h.bookService.UpdateBookRating(bookID, req.Rating)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	err := h.bookService.DeleteBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *Handler) GetBook(c *gin.Context) {
	bookID := c.Param("id")

	book, err := h.bookService.GetBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}

// User handlers

func (h *Handler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	user, err := h.userService.CreateUser(req.Name, req.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserToResponse(user)
	c.JSON(201, response)
}

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.userService.GetUser(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserToResponse(user)
	c.JSON(200, response)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := h.userService.DeleteUser(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get users"})
		return
	}

	response := dto.UsersToResponse(users)
	c.JSON(200, response)
}
