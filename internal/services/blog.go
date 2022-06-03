package services

import (
	"context"
	"github.com/Loner1024/uniix.io/api/gen/go/api"
	"github.com/Loner1024/uniix.io/internal/aggregate"
	"github.com/Loner1024/uniix.io/internal/domain/blog"
	"github.com/Loner1024/uniix.io/internal/entity"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	useCase *blog.UseCase
}

// NewService creates a new service
func NewService(userCase *blog.UseCase) *Service {
	return &Service{useCase: userCase}
}

// GetArticle get an article by id.
func (s Service) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (*v1.Article, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}
	data, err := s.useCase.GetBlog(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, "get article error")
	}
	return blogToDTO(data), nil
}

// ListArticle list article.
func (s Service) ListArticle(ctx context.Context, req *v1.ListArticleRequest) (*v1.ListArticleResponse, error) {
	data, err := s.useCase.ListBlog(ctx, req.GetPageSize(), req.GetPageToken())
	if err != nil {
		return nil, status.Error(codes.Internal, "list article error")
	}
	
	return &v1.ListArticleResponse{
		Articles:      blogSliceToDTO(data.Blogs),
		NextPageToken: data.NextPageToken,
	}, nil
}

func blogSliceToDTO(data []aggregate.Blog) []*v1.Article {
	result := make([]*v1.Article, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = blogToDTO(data[i])
	}
	return result
}

func blogToDTO(data aggregate.Blog) *v1.Article {
	return &v1.Article{
		Id:         data.ID.String(),
		Title:      data.Title,
		Content:    data.Content,
		CreateTime: timestamppb.New(data.CreateTime),
		UpdateTime: timestamppb.New(data.UpdateTime),
		Tags:       tagSliceToDTO(data.Tags),
	}
}

func tagSliceToDTO(data []entity.Tag) []*v1.Tag {
	tagsData := make([]*v1.Tag, len(data))
	for i := 0; i < len(data); i++ {
		tagsData[i] = tagToDTO(data[i])
	}
	return tagsData
}

func tagToDTO(data entity.Tag) *v1.Tag {
	return &v1.Tag{
		Value: data.Value,
	}
}
