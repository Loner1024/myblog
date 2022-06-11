package firebase

import (
	"github.com/Loner1024/uniix.io/internal/entity"
	"github.com/google/uuid"
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
