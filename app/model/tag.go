package model

type Tag struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Status    uint64 `json:"status"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
	DeletedAt uint64 `json:"deleted_at"`
}

type CreateTagRequest struct {
	Name string `json:"name"`
}

type CreateTagResponse struct {
	Tag Tag `json:"tag"`
}

type ListTagResponse struct {
	Tags []Tag `json:"tags"`
}
