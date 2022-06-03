package entity

import "testing"

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestNewArticle(t *testing.T) {
	_, err := NewArticle("title", "content")
	if err != nil {
		t.Fatalf("\tTest %s:\tWhen create a new article entity.", Failed)
	}
}
