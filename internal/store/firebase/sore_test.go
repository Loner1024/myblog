package firebase

import (
	"context"
	"github.com/Loner1024/uniix.io/configs"
	"github.com/Loner1024/uniix.io/internal/aggregate"
	"github.com/Loner1024/uniix.io/internal/entity"
	"github.com/google/uuid"
	"runtime"
	"testing"
	"time"
)

var conf = configs.Config{
	Firebase: configs.Firebase{AgentFile: "../../../configs/firebase-agent.json"},
}

func TestCreateBlog(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("skip local test code")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if err := ctx.Err(); err != nil {
			t.Errorf("ctx timeout: %v", err)
		}
		cancel()
	}()
	store, err := NewStore(conf)
	if err != nil {
		t.Fatal(err)
	}
	err = store.CreateBlog(ctx, aggregate.Blog{
		Article: entity.Article{
			ID:         uuid.New(),
			Title:      "这里是文章标题",
			Content:    "内容",
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
		Tags: []entity.Tag{
			{
				Value: "标签1",
			},
			{
				Value: "标签2",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBlog(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("skip local test code")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if err := ctx.Err(); err != nil {
			t.Errorf("ctx timeout: %v", err)
		}
		cancel()
	}()
	store, err := NewStore(conf)
	if err != nil {
		t.Fatal(err)
	}
	blog, err := store.GetBlog(ctx, uuid.MustParse("018045aa-8949-4a61-986d-f1935e453001"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", blog)
}
