package model

type PostTag struct {
	PostId uint64 `json:"post_id"`
	TagId  uint64 `json:"tag_id"`
}
