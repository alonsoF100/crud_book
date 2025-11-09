package models

type Book struct {
	ID          string  `json:"id" db:"id"`
	UserID      string  `json:"user_id" db:"user_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Rating      float64 `json:"rating" db:"rating"`
	Status      string  `json:"status" db:"status"`
}

type User struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}
