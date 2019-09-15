package blog

import (
	"github.com/lexkong/log"
	"go_simpleweibo/database"
)

// 对blog用户模型进行增删改

func (b *Blog) Create() (err error) {
	if err = database.DB.Create(&b).Error; err != nil {
		log.Warnf("微博创建失败: %v", err)
		return err
	}

	return nil
}

func  Delete(id int) (err error) {
	blog := &Blog{}
	blog.BaseModel.ID = uint(id)

	if err = database.DB.Delete(&blog).Error; err != nil {
		log.Warnf("微博删除失败：%v", err)
		return err
	}
	return nil
}
