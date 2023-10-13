package model

type User struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	AttachmentId string `json:"attachment_id"`
	Status       uint64 `json:"status"`
	CreatedAt    uint64 `json:"created_at"`
	UpdatedAt    uint64 `json:"updated_at"`
	DeletedAt    uint64 `json:"deleted_at"`
}

type CreateUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	AttachmentId string `json:"attachment_id"`
}

type CreateUserResponse struct {
	User User `json:"user"`
}

type ListUserResponse struct {
	Users []User `json:"users"`
}
