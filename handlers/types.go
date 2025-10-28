package handlers

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateBookRequest struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateBookStatusRequest struct {
	Status string `json:"status"`
}

type UpdateBookRatingRequest struct {
	Rating float64 `json:"rating"`
}
