package models

type StreamModel[T interface{}] struct {
	Model     T    `json:"model"`
	IsDeleted bool `json:"is_deleted"`
}
