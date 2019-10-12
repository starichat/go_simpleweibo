package models

import "time"

const (
	// TrueTinyint true
	TrueTinyint = 1
	// FalseTinyint false
	FalseTinyint = 0
)

// BaseModel model 基类
type BaseModel struct {
	ID uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	// MySQL 的 DATE/DATETIME类型可以对应Goland的time.Time
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// 有 DeletedAt(类型需要是 *time.Time) 即支持 gorm 软删除
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index"`
}
