package handlers

import "crud_book/internal/services"

type handlers struct {
	bookHandler *bookHandler
	userHandler *userHandler
}

func NewHandlers(bookService services.BookService, userService services.UserService) *handlers {
	return &handlers{
		bookHandler: newBookHandler(bookService),
		userHandler: newUserHandler(userService),
	}
}
