package blog

import "go_simpleweibo/app/models"

// blog 博客内容 模型
type Blog struct {
	models.BaseModel
	Content string `gorm:"column:context;type:text;not null"`
	UserID  uint   `gorm:"column:user_id;not null" sql:"index"` // 一对多，关联 User Model
}

func (Blog) TableName() string {
	return "blogs"
}
