package entity

import (
	"github.com/google/uuid"
	"time"
)

// Article is article entity.
type Article struct {
	ID         uuid.UUID
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
}
