package password


import (
	"go_simpleweibo/routes/named"
	"github.com/gin-gonic/gin"
	"go_simpleweibo/app/helpers"
	passwordResetModel "go_simpleweibo/app/models/password_reset"
)

func sendResetEmail(pwd *passwordResetModel.PasswordReset) error {
	subject := "重置密码！请确认你的邮箱。"
	tpl := "mail/reset_password.html"
	resetPasswordURL := named.G("password.reset", "token", pwd.Token)

	return helpers.SendMail([]string{pwd.Email}, subject, tpl, gin.H{"resetPasswordURL": resetPasswordURL})

}