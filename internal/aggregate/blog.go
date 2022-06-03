package aggregate

import "github.com/Loner1024/uniix.io/internal/entity"

type Blog struct {
	entity.Article
	Tags []entity.Tag
}
