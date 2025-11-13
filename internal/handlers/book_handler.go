package handlers

import (
	"context"
	"crud_book/internal/dto"
	"crud_book/internal/services"
	"time"

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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	book, err := h.bookService.CreateBook(ctx, req)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	books, err := h.bookService.GetUserBooks(ctx, req)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.BooksToResponse(books)
	c.JSON(200, response)
}

func (h *bookHandler) updateBookStatus(c *gin.Context) {
	var req dto.UpdateBookStatusRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	book, err := h.bookService.UpdateBookStatus(ctx, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}

func (h *bookHandler) updateBookRating(c *gin.Context) {
	var req dto.UpdateBookRatingRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	book, err := h.bookService.UpdateBookRating(ctx, req)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	err := h.bookService.DeleteBook(ctx, req)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	book, err := h.bookService.GetBook(ctx, req)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}
