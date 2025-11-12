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

	book, err := h.bookService.CreateBook(req.UserID, req.Name, req.Description)
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
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

func (h *bookHandler) updateBookRating(c *gin.Context) {
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

func (h *bookHandler) deleteBook(c *gin.Context) {
	bookID := c.Param("id")

	err := h.bookService.DeleteBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *bookHandler) getBook(c *gin.Context) {
	bookID := c.Param("id")

	book, err := h.bookService.GetBook(bookID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.BookToResponse(book)
	c.JSON(200, response)
}
