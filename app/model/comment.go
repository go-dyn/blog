package model

type Comment struct {
	ID        uint64 `json:"id"`
	UserId    uint64 `json:"user_id"`
	PostId    uint64 `json:"post_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Status    uint64 `json:"status"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
	DeletedAt uint64 `json:"deleted_at"`
}

type CreateCommentRequest struct {
	UserId string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type CreateCommentResponse struct {
	Comment Comment `json:"comment"`
}

type ListCommentResponse struct {
	Comments []Comment `json:"comments"`
}
