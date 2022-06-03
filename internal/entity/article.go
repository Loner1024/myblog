package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrArticleNoTitle = errors.New("article must need a title")
)

// Article is article entity.
type Article struct {
	ID         uuid.UUID
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
}

// NewArticle create an article entity.
func NewArticle(title, content string) (Article, error) {
	if title == "" {
		return Article{}, ErrArticleNoTitle
	}
	return Article{
		ID:         uuid.New(),
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
	}, nil
}
