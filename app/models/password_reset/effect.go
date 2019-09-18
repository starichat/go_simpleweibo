package passwordreset

import (
	"go_simpleweibo/database"
	"github.com/lexkong/log"
	"go_simpleweibo/pkg/utils"
)


func (p *PasswordReset) Create() (err error) {
	token := string(utils.RandomCreateBytes(30))

	// 如果已经存在则删除
	if existedPwd, err := GetByEmail(p.Email); err == nil {
		if err = DeleteByEmail(existedPwd.Email); err != nil {
			return err
		}
	}

	// 创建
	p.Token= token
	if err = database.DB.Create(&p).Error; err != nil {
		 log.Warnf("%s 创建失败：%v",p.TableName(), err)
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



