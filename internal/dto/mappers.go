package dto

import "crud_book/internal/models"

func BookToResponse(book *models.Book) BookResponse {
	return BookResponse{
		ID:          book.ID,
		UserID:      book.UserID,
		Name:        book.Name,
		Description: book.Description,
		Rating:      book.Rating,
		Status:      book.Status,
	}
}

func UserToResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func BooksToResponse(books []*models.Book) []BookResponse {
	responses := make([]BookResponse, len(books))
	for i, book := range books {
		responses[i] = BookToResponse(book)
	}
	return responses
}

func UsersToResponse(users []*models.User) []UserResponse {
	responses := make([]UserResponse, len(users))
	for i, user := range users {
		responses[i] = UserToResponse(user)
	}
	return responses
}
