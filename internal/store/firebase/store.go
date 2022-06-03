package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"github.com/Loner1024/uniix.io/configs"
	"github.com/Loner1024/uniix.io/internal/aggregate"
	"github.com/Loner1024/uniix.io/internal/domain/blog"
	"github.com/Loner1024/uniix.io/internal/entity"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"strings"
	"time"
)

type Store struct {
	store *firestore.Client
}

const (
	ArticleCollection = "articles"
	TagCollection     = "tags"
)

type ArticleDoc struct {
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
	Tags       []*firestore.DocumentRef
}

type TagDoc struct {
	Article []*firestore.DocumentRef
}

func (s Store) GetBlog(ctx context.Context, ID uuid.UUID) (aggregate.Blog, error) {
	dsnap, err := s.store.Collection(ArticleCollection).Doc(ID.String()).Get(ctx)
	if err != nil {
		return aggregate.Blog{}, err
	}
	var articleDoc ArticleDoc
	err = dsnap.DataTo(&articleDoc)
	if err != nil {
		return aggregate.Blog{}, err
	}
	tags := make([]entity.Tag, len(articleDoc.Tags))
	for i := 0; i < len(tags); i++ {
		paths := strings.Split(articleDoc.Tags[i].Path, "/")
		tags[i] = entity.Tag{Value: paths[len(paths)-1]}
	}
	return aggregate.Blog{
		Article: entity.Article{
			ID:         ID,
			Title:      articleDoc.Title,
			Content:    articleDoc.Content,
			CreateTime: articleDoc.CreateTime,
			UpdateTime: articleDoc.UpdateTime,
		},
		Tags: tags,
	}, nil
}

func (s Store) ListBlog(ctx context.Context, limit, offset int64) ([]aggregate.Blog, error) {
	// TODO implement me
	panic("implement me")
}

func (s Store) CountBlog(ctx context.Context) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (s Store) CreateBlog(ctx context.Context, data aggregate.Blog) error {
	err := s.store.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		tagRefs := make([]*firestore.DocumentRef, len(data.Tags))
		for i := 0; i < len(data.Tags); i++ {
			tagRefs[i] = s.store.Collection(TagCollection).Doc(data.Tags[i].Value)
		}
		docRef := s.store.Collection(ArticleCollection).Doc(data.ID.String())
		if err := tx.Create(docRef, ArticleDoc{
			Title:      data.Title,
			Content:    data.Content,
			CreateTime: data.CreateTime,
			UpdateTime: data.UpdateTime,
			Tags:       tagRefs,
		}); err != nil {
			return err
		}
		for i := 0; i < len(tagRefs); i++ {
			if err := tx.Set(tagRefs[i], map[string]interface{}{
				"article_ref": firestore.ArrayUnion(docRef),
			}, firestore.MergeAll); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func NewStore(conf configs.Config) (blog.Repo, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(conf.Firebase.AgentFile))
	if err != nil {
		return nil, err
	}
	store, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return &Store{store: store}, nil
}
