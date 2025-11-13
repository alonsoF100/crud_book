package handlers

import (
	"crud_book/internal/dto"
	"crud_book/internal/services"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService services.BookService
}

func newBookHandler(bookService services.BookService) *bookHandler {
	return &bookHandler{bookService: bookService}
}

func (h *bookHandler) createBook(c *gin.Context) {
	var req dto.CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	book, err := h.bookService.CreateBook(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(201, response)
}

func (h *bookHandler) getUserBooks(c *gin.Context) {
	var req dto.GetUserBooksRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.SetDefaults()

	books, err := h.bookService.GetUserBooks(req)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.BooksToResponse(books)
	c.JSON(200, response)
}

func (h *bookHandler) updateBookStatus(c *gin.Context) {
	var req dto.UpdateBookStatusRequest

	// 1. Сначала биндим URI (BookID)
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 2. Биндим JSON (Status)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.UpdateBookStatus(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}

func (h *bookHandler) updateBookRating(c *gin.Context) {
	var req dto.UpdateBookRatingRequest

	// 1. Сначала биндим URI (BookID)
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 2. Биндим JSON (Status)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.UpdateBookRating(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}

func (h *bookHandler) deleteBook(c *gin.Context) {
	var req dto.DeleteBookRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}

	err := h.bookService.DeleteBook(req)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *bookHandler) getBook(c *gin.Context) {
	var req dto.GetBookRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.bookService.GetBook(req)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}
