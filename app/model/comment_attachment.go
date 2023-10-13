package model

type CommentAttachment struct {
	CommentId    uint64 `json:"comment_id"`
	AttachmentId uint64 `json:"attachment_id"`
}
