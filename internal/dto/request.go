package dto

// Book Requests
type CreateBookRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateBookStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=want reading finished"`
}

type UpdateBookRatingRequest struct {
	Rating float64 `json:"rating" binding:"required,min=0,max=5"`
}

type GetUserBooksRequest struct {
	UserID    string  `form:"-" uri:"id" binding:"required"`
	Limit     int     `form:"limit" binding:"omitempty,min=1,max=50"`
	Offset    int     `form:"offset" binding:"omitempty,min=0"`
	Sort      string  `form:"sort" binding:"omitempty,oneof=created_at name rating status"`
	Order     string  `form:"order" binding:"omitempty,oneof=asc desc"`
	Status    string  `form:"status" binding:"omitempty,oneof=want reading finished"`
	MinRating float64 `form:"min_rating" binding:"omitempty,min=0,max=5"`
	MaxRating float64 `form:"max_rating" binding:"omitempty,min=0,max=5"`
}

func (r *GetUserBooksRequest) SetDefaults() {
	if r.Limit == 0 {
		r.Limit = 20
	}
	if r.Sort == "" {
		r.Sort = "created_at"
	}
	if r.Order == "" {
		r.Order = "desc"
	}
}

// User Requests
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
