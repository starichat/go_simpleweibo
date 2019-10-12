package blog

import (
	"go_simpleweibo/app/models"
	"go_simpleweibo/database"

	"github.com/lexkong/log"
)

const tableName = "blogs"

// Status 微博
type Blog struct {
	models.BaseModel
	Content string `gorm:"column:context;type:text;not null"`
	UserID  uint   `gorm:"column:user_id;not null" sql:"index"` // 一对多，关联 User Model
}

// TableName 表名
func (Blog) TableName() string {
	return tableName
}

// Create -
func (b *Blog) Create() (err error) {
	if err = database.DB.Create(&b).Error; err != nil {
		log.Warnf("微博创建失败: %v", err)
		return err
	}

	return nil
}

// Delete -
func Delete(id int) (err error) {
	blog := &Blog{}
	blog.BaseModel.ID = uint(id)

	if err = database.DB.Delete(&blog).Error; err != nil {
		log.Warnf("微博删除失败: %v", err)
		return err
	}

	return nil
}
