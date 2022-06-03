package firebase

import (
	"context"
	"github.com/Loner1024/uniix.io/internal/aggregate"
	"github.com/Loner1024/uniix.io/internal/domain/blog"
	"github.com/google/uuid"
)

type Store struct{}

func (s Store) GetBlog(ctx context.Context, ID uuid.UUID) (aggregate.Blog, error) {
	// TODO implement me
	panic("implement me")
}

func (s Store) ListBlog(ctx context.Context, limit, offset int64) ([]aggregate.Blog, error) {
	// TODO implement me
	panic("implement me")
}

func (s Store) CountBlog(ctx context.Context) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func NewStore() blog.Repo {
	return &Store{}
}
