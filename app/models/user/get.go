package user

import (
	"crypto/md5"
	"encoding/hex"
	"go_simpleweibo/app/models"
	"go_simpleweibo/database"
	"strconv"
)

// GetById
func Get(id int) (*User, error) {
	user := &User{}
	d := database.DB.First(&user, id)
	return user, d.Error
}

// 通过邮箱查询用户信息
func GetByEmail(email string) (*User, error) {
	user := &User{}
	d := database.DB.Where("email = ?", email).First(&user)
	return user, d.Error
}

// 通过激活token查询用户信息
func GetByActivationToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("activation_token = ?", token).First(&user)
	return user, d.Error
}

// 通过remembertoken
func GetByRememberToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("remember_token = ?", token).First(&user)
	return user, d.Error
}

// 获取用户列表（分页查询）
func GetUserList(offset, limit int) (users []*User, err error) {
	users = make([]*User, 0)

	if err := database.DB.Offset(offset).Limit(limit).Order("id").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// 总用户数
func GetAllCount() (count int, err error) {
	err = database.DB.Model(&User{}).Count(&count).Error
	return
}

// 获取用户头像
func (u *User) GetAvatar() string {
	if u.Avatar != "" {
		return u.Avatar
	}
	hash := md5.Sum([]byte(u.Email))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
}

// GetIDstring 获取字符串形式的 id
func (u *User) GetIDstring() string {
	return strconv.Itoa(int(u.ID))
}

// IsAdminRole 是否为管理员
func (u *User) IsAdminRole() bool {
	return u.IsAdmin == models.TrueTinyint
}

// IsActivated 是否已激活
func (u *User) IsActivated() bool {
	return u.Activated == models.TrueTinyint
}
