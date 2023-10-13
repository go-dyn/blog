package model

type Attachment struct {
	ID        uint64 `json:"id"`
	UserId    uint64 `json:"user_id"`
	Name      string `json:"name"`
	SrcName   string `json:"src_name"`
	Ext       string `json:"ext"`
	Path      string `json:"path"`
	Mime      string `json:"mime"`
	Size      string `json:"size"`
	Type      string `json:"type"`
	Status    uint64 `json:"status"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
	DeletedAt uint64 `json:"deleted_at"`
}

type CreateAttachmentRequest struct {
	UserId  uint64 `json:"user_id"`
	Name    string `json:"name"`
	SrcName string `json:"src_name"`
	Ext     string `json:"ext"`
	Path    string `json:"path"`
	Mime    string `json:"mime"`
	Size    string `json:"size"`
	Type    string `json:"type"`
	Status  uint64 `json:"status"`
}

type CreateAttachmentResponse struct {
	Attachment Attachment `json:"attachment"`
}

type ListAttachmentResponse struct {
	Attachments []Attachment `json:"attachments"`
}
