package blog

import (
	"context"
	"fmt"
	"github.com/Loner1024/uniix.io/internal/aggregate"
	"github.com/Loner1024/uniix.io/utils/pagination"
	"github.com/google/uuid"
)

type UseCase struct {
	repo Repo
}

func NewUseCase(repo Repo) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) GetBlog(ctx context.Context, id uuid.UUID) (aggregate.Blog, error) {
	return u.repo.GetBlog(ctx, id)
}

func (u *UseCase) ListBlog(ctx context.Context, pageSize int32, pageToken string) (aggregate.BlogList, error) {
	page, err := pagination.ParsePageToken(pageSize, pageToken)
	if err != nil {
		return aggregate.BlogList{}, fmt.Errorf("parse page token:%w", err)
	}
	data, err := u.repo.ListBlog(ctx, int64(pageSize), page.Offset)
	if err != nil {
		return aggregate.BlogList{}, err
	}
	
	return aggregate.BlogList{
		Blogs:         data,
		NextPageToken: page.Next(pageSize).String(),
	}, nil
}
