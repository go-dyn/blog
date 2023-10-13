package model

type Post struct {
	ID        uint64 `json:"id"`
	UserId    uint64 `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Status    uint64 `json:"status"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
	DeletedAt uint64 `json:"deleted_at"`
}

type CreatePostRequest struct {
	UserId string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type CreatePostResponse struct {
	Post Post `json:"post"`
}

type ListPostResponse struct {
	Posts []Post `json:"posts"`
}
