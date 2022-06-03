package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTag(t *testing.T) {
	tagValue := "tag"
	tag, err := NewTag(tagValue)
	assert.NoError(t, err, ErrTagEmpty)
	assert.Equal(t, tagValue, tag.Value)
}

func TestEmptyTag(t *testing.T) {
	_, err := NewTag("")
	assert.ErrorIs(t, err, ErrTagEmpty)
}
