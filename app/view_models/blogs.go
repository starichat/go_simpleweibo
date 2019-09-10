package view_models

import (
	blogModel "go_simpleweibo/app/models/blog"
	"go_simpleweibo/pkg/time"
)

type BlogViewModel struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt string
}

// NewBlogsViewModelSerializer 微博数据展示
func NewBlogsViewModelSerializer(s *blogModel.Blog) *BlogViewModel {
	return &BlogViewModel{
		ID:        int(s.ID),
		Content:   s.Content,
		UserID:    int(s.UserID),
		CreatedAt: time.SinceForHuman(s.CreatedAt),
	}
}
