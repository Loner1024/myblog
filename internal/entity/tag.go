package entity

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrTagEmpty = errors.New("tag not allow empty")
)

type Tag struct {
	ID  uuid.UUID
	Tag string
}

func NewTag(tag string) (Tag, error) {
	if tag == "" {
		return Tag{}, ErrTagEmpty
	}
	return Tag{
		ID:  uuid.New(),
		Tag: tag,
	}, nil
}
