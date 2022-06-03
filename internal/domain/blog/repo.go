package blog

import (
	"context"
	"github.com/Loner1024/uniix.io/internal/aggregate"
	"github.com/google/uuid"
)

type Repo interface {
	GetBlog(ctx context.Context, ID uuid.UUID) (aggregate.Blog, error)
	ListBlog(ctx context.Context, limit, offset int64) ([]aggregate.Blog, error)
	CreateBlog(ctx context.Context, data aggregate.Blog) error
	CountBlog(ctx context.Context) (int64, error)
}
