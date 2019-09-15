package blog

import (
	"github.com/gin-gonic/gin"
	"go_simpleweibo/app/controllers"
	blogModel "go_simpleweibo/app/models/blog"
	userModel "go_simpleweibo/app/models/user"
	"go_simpleweibo/pkg/flash"
)

// 创建微博
func Store(c *gin.Context,currentUser *userModel.User) {
	// 从表格中获取微博内容
	content:=c.DefaultPostForm("content","") // 没有获取到content的值就给一个默认值
	// 截取微博内容长度，一便对不合规微博进行判断同时给出提示信息
	contentLenth:=len(content)
	// 对长度进行错误处理
	if contentLenth==0{
		flash.NewDangerFlash(c,"微博内容不能为空")
		backTo(c,currentUser)
		return
		}
	if contentLenth > 140 {
		flash.NewDangerFlash(c,"微博内容不能超过140词")
		backTo(c, currentUser)
		return
	}

	// 生成微博实例模型
	blog:=&blogModel.Blog{
		Content:content,
		UserID:currentUser.ID,
	}
	// 生成数据库并发布
	if err := blog.Create();err != nil {
		flash.NewDangerFlash(c,"发布失败")
		backTo(c,currentUser)
		return
	}
	flash.NewSuccessFlash(c,"发布成功")
	backTo(c, currentUser)
}

// 删除微博
func Destroy(c *gin.Context,currentUser *userModel.User) {
	// 根据path获取id参数
	blogID, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}
	// 根据blogID 查询blog
	blog, err := blogModel.Get(blogID)
	if err != nil {
		flash.NewDangerFlash(c, "删除失败")
		backTo(c, currentUser)
		return
	}

	// 权限判断

	// 删除微博
	if err := blogModel.Delete(int(blog.ID)); err != nil {
		flash.NewDangerFlash(c, "删除失败")
		backTo(c, currentUser)
		return
	}

	flash.NewSuccessFlash(c, "删除成功")
	backTo(c, currentUser)
}
