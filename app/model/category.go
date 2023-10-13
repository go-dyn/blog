package model

type Category struct {
	ID        uint64 `json:"id"`
	ParentId  uint64 `json:"parent_id"`
	Name      string `json:"name"`
	Status    uint64 `json:"status"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
	DeletedAt uint64 `json:"deleted_at"`
}

type CreateCategoryRequest struct {
	UserId string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type CreateCategoryResponse struct {
	Category Category `json:"category"`
}

type ListCategoryResponse struct {
	Posts []Post `json:"categories"`
}
