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
	tags := make([]entity.Tag, 0, len(articleDoc.Tags)+1)
	tagDocRefToEntityTag(articleDoc.Tags, &tags)
	return aggregate.Blog{
		Article: articlDocToEntityArticle(ID, articleDoc),
		Tags:    tags,
	}, nil
}

func (s Store) ListBlog(ctx context.Context, limit, offset int64) ([]aggregate.Blog, error) {
	docs := s.store.Collection(ArticleCollection).
		Select("Title", "UpdateTime", "Tags", "CreateTime").
		Limit(int(limit)).
		Offset(int(offset)).
		OrderBy("CreateTime", firestore.Desc).
		Documents(ctx)
	
	doc, err := docs.GetAll()
	if err != nil {
		return nil, err
	}
	result := make([]aggregate.Blog, 0, limit+1)
	for _, v := range doc {
		var articleDoc ArticleDoc
		err = v.DataTo(&articleDoc)
		if err != nil {
			return nil, err
		}
		tags := make([]entity.Tag, 0, len(articleDoc.Tags)+1)
		tagDocRefToEntityTag(articleDoc.Tags, &tags)
		result = append(result, aggregate.Blog{
			Article: articlDocToEntityArticle(uuid.MustParse(v.Ref.ID), articleDoc),
			Tags:    tags,
		})
	}
	return result, nil
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
