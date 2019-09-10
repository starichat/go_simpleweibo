package password_reset

import "time"

// passwordReset 重置密码模型
type PasswordReset struct {
	Email     string    `gorm:"column:email;type:varchar(255);not null" sql:"index"`
	Token     string    `gorm:"column:token;type:varchar(255);not null" sql:"index"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (PasswordReset) TableName() string {
	return "password_resets"
}
