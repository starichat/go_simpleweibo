package user

import (
	"go_simpleweibo/app/models"
	"go_simpleweibo/database"
	"go_simpleweibo/pkg/utils"
	"time"

	"github.com/lexkong/log"
)

// 用户模型
type User struct {
	models.BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	Email    string `gorm:"column:email;type:varchar(255);unique;not null"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
	// 是否为管理员
	IsAdmin uint `gorm:"column:is_admin;type:tinyint(1)"`
	// 用户激活
	ActivationToken string    `gorm:"column:activation_token;type:varchar(255)"`
	Activated       uint      `gorm:"column:activated;type:tinyint(1);not null"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"` // 激活时间
	// 用于实现记住我功能，存入 cookie 中，下次带上时，即可直接登录
	RememberToken string `gorm:"column:remember_token;type:varchar(100)"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) CreateUser() (err error) {
	//TODO : 加密密码
	// 生成用户 remember_token
	if u.RememberToken == "" {
		u.RememberToken = string(utils.RandomCreateBytes(10))
	}
	// 生成用户激活 token
	if u.ActivationToken == "" {
		u.ActivationToken = string(utils.RandomCreateBytes(30))
	}
	if err = database.DB.Create(&u).Error; err != nil {
		log.Warnf("user create failed:%v", err)
		return err
	}
	return nil

}

// 更新用户
func (u *User) Update(needEncryotPwd bool) (err error) {
	if needEncryotPwd {
		// TODO：更新加密
	}

	if err = database.DB.Save(&u).Error; err != nil {
		log.Warnf("用户更新失败:&v", err)
		return err
	}

	return nil
}

// Delete -
func Delete(id int) (err error) {
	user := &User{}
	user.BaseModel.ID = uint(id)

	// Unscoped: 永久删除而不是软删除 (由于该操作是管理员操作的，所以不使用软删除)
	if err = database.DB.Unscoped().Delete(&user).Error; err != nil {
		log.Warnf("用户删除失败: %v", err)
		return err
	}

	return nil
}
