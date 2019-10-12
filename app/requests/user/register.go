package user

// 执行验证请求的功能

// 以后可以改为 tag 来调用验证器函数
type UserCreateForm struct {
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"` 
}

// 验证邮箱是否被注册
// func (u *UserCreateForm) emailUniqueValidator() requests.ValidatorFunc {
// 	return func() (msg string) {
// 		if _, err := userModel.GetByEmail(u.Email); err != nil {
// 			return ""
// 		}
// 		return "邮箱已经被注册过了"
// 	}
// }

// 验证用户传递参数是否合法
func (u *UserCreateForm) Validate() (errors []string) {

}

// 验证成功并创建用户
func (u *UserCreateForm) ValidateAndSave() (user *userModel.User, errors []string) {

	return 
}