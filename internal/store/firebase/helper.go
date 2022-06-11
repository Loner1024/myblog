package firebase

import (
	"cloud.google.com/go/firestore"
	"github.com/Loner1024/uniix.io/internal/entity"
	"github.com/google/uuid"
	"strings"
)

func articlDocToEntityArticle(ID uuid.UUID, doc ArticleDoc) entity.Article {
	return entity.Article{
		ID:         ID,
		Title:      doc.Title,
		Content:    doc.Content,
		CreateTime: doc.CreateTime,
		UpdateTime: doc.UpdateTime,
	}
}

func tagDocRefToEntityTag(tags []*firestore.DocumentRef, result *[]entity.Tag) {
	for i := 0; i < len(tags); i++ {
		paths := strings.Split(tags[i].Path, "/")
		*result = append(*result, entity.Tag{Value: paths[len(paths)-1]})
	}
}
