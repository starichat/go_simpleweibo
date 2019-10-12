package user

import (
	"github.com/gin-gonic/gin"
	"go_simpleweibo/pkg/response"
	"net/http"
	"time"
	userRequest "go_simpleweibo/app/requests/user"
)

func Register(c *gin.Context) {
	//1. 验证参数和创建用户
	userCreateForm := &userRequest.UserCreateForm{
		Name:                 c.PostForm("name"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}

	//2. 数据绑定和校验
	if err := c.ShouldBindJSON(&userCreateForm); err != nil {
		c.JSON(http.StatusBadRequest, response.Error{Error:err}
		return 
	}

	//3. 验证参数并获取结果
	user, error := userRequest.ValidateAndSave()

	//4. 判断error情况,有错误则返回错误，进行错误处理，给用户一个合适的提示
	if len(error) != 0 || user == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error:error})
		return 
	}

	//5. 发送注册邮件
	if err := sendConfirmEmail(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error:邮箱发送失败})
		return 
	}

	c.JSON(http.StatusRequest, response.Success{Success:"验证邮件已发送到你的注册邮箱上，请注意查收。"})
	
}



// 确认邮箱：确认成功后执行持久化用户数据
func ConfirmEmail(c *gin.Context){
	token := c.Param("token")

	// 根据激活token 查询用户信息
	user, err := userModel.GetByActivationToken(token)
	if user == nil || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error:""})
		return
	}

	// 更新用户
	user.Activated = model.TrueTinyint
	user.ActivationToken = ""  // 将token值为空 
	user.EmailVerifiedAt = time.Now()
	if err = user.Update(false); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error:""})
		return 
	}

	// 用当前用户信息执行登录，并跳转登录界面,并提示登录成功信息
	c.JSON(http.StatusOK,"激活成功")

}

// 用户登录
func Login(c *gin.Context){
	// 获取json数据
	// 校验数据是否正确
	// 错误分析执行登录，返回token
}

// 用户注销
func Logout(c *gin.Context){
	// 清除token
}

