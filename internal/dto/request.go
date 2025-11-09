package dto

// Book Requests
type CreateBookRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateBookStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=reading finished planned"`
}

type UpdateBookRatingRequest struct {
	Rating float64 `json:"rating" binding:"required,min=0,max=5"`
}

// User Requests
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
