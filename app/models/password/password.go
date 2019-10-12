package password

import (
	"go_simpleweibo/database"
	"go_simpleweibo/pkg/utils"
	"time"

	"github.com/lexkong/log"
)

type PasswordReset struct {
	Email     string    `gorm:"column:email;type:varchar(255);not null" sql:"index"`
	Token     string    `gorm:"column:token;type:varchar(255);not null" sql:"index"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (PasswordReset) TableName() string {
	return "password_resets"
}

// Create -
func (p *PasswordReset) Create() (err error) {
	token := string(utils.RandomCreateBytes(30))

	// 如已存在则先删除 (可以判断下，不能创建太频繁)
	if existedPwd, err := GetByEmail(p.Email); err == nil {
		if err = DeleteByEmail(existedPwd.Email); err != nil {
			return err
		}
	}

	// 创建
	p.Token = token
	if err = database.DB.Create(&p).Error; err != nil {
		log.Warnf("%s 创建失败: %v", p.TableName(), err)
		return err
	}

	return nil
}

// DeleteByEmail -
func DeleteByEmail(email string) (err error) {
	pwd := &PasswordReset{}

	if err = database.DB.Where("email = ?", email).Delete(&pwd).Error; err != nil {
		log.Warnf("%s 删除失败: %v", pwd.TableName(), err)
		return err
	}

	return nil
}

// DeleteByToken -
func DeleteByToken(token string) (err error) {
	pwd := &PasswordReset{}

	if err = database.DB.Where("token = ?", token).Delete(&pwd).Error; err != nil {
		log.Warnf("%s 删除失败: %v", pwd.TableName(), err)
		return err
	}

	return nil
}
