package entity

import (
	"errors"
)

var (
	ErrTagEmpty = errors.New("tag not allow empty")
)

type Tag struct {
	Value string
}

func NewTag(tag string) (Tag, error) {
	if tag == "" {
		return Tag{}, ErrTagEmpty
	}
	return Tag{
		Value: tag,
	}, nil
}
