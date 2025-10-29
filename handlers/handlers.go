package handlers

import (
	"crud_book/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Handler {
	return &Handler{storage: storage}
}

// Хендлеры для book

func (h *Handler) CreateBook(c *gin.Context) {
	type CreateBookRequest struct {
		UserID      string `json:"user_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var req CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Неверный JSON"})
		return
	}

	book, err := h.storage.AddBook(req.UserID, req.Name, req.Description)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, book)
}

func (h *Handler) GetUserBooks(c *gin.Context) {
	userID := c.Param("id")

	books, err := h.storage.GetUserBooks(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, books)
}

func (h *Handler) UpdateBookStatus(c *gin.Context) {
	type UpdateBookStatusRequest struct {
		Status string `json:"status"`
	}
	var req UpdateBookStatusRequest

	bookID := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Неверный JSON"})
		return
	}

	err := h.storage.UpdateBookStatus(bookID, req.Status)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book, err := h.storage.GetBook(bookID)
	if err != nil {
		c.JSON(500, gin.H{"error": "не удалось получить обновленную книгу"})
		return
	}
	c.JSON(200, book)
}

func (h *Handler) UpdateBookRating(c *gin.Context) {
	type UpdateBookRatingRequest struct {
		Rating float64 `json:"rating"`
	}
	var req UpdateBookRatingRequest

	bookID := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Неверный JSON"})
		return
	}

	err := h.storage.UpdateBookRating(bookID, req.Rating)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book, err := h.storage.GetBook(bookID)
	if err != nil {
		c.JSON(500, gin.H{"error": "не удалось получить обновленную книгу"})
		return
	}
	c.JSON(200, book)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	err := h.storage.DeleteBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *Handler) GetBook(c *gin.Context) {
	bookID := c.Param("id")

	book, err := h.storage.GetBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, book)
}

// Хендлеры для user

func (h *Handler) CreateUser(c *gin.Context) {
	type CreateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Неверный JSON"})
		return
	}

	user, err := h.storage.AddUser(req.Name, req.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.storage.GetUser(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := h.storage.DeleteUser(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users := h.storage.GetAllUsers()

	c.JSON(200, users)
}
