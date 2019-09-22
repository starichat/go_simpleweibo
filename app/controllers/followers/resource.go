package followers

import (
	"fmt"
	"go_simpleweibo/app/controllers"
	followerModel "go_simpleweibo/app/models/follower"
	userModel "go_simpleweibo/app/models/user"
	"go_simpleweibo/pkg/flash"
	"go_simpleweibo/routes/named"

	"github.com/gin-gonic/gin"
)

// Store 关注用户
func Store(c *gin.Context, currentUser *userModel.User) {
	id, err := controllers.GetIntParam(c, "id")
	fmt.Println("id",id)
	if err != nil {
		controllers.Render404(c)
		return
	}
	isFollowing := false
	if id != int(currentUser.ID) {
		fmt.Println("1111",int(currentUser.ID))
		isFollowing = followerModel.IsFollowing(int(currentUser.ID), id)
		fmt.Println(isFollowing)
	}

	if !isFollowing {
		fmt.Println(currentUser.ID)
		fmt.Println(uint(id))
		if err := followerModel.DoFollow(currentUser.ID, uint(id)); err != nil {
			flash.NewDangerFlash(c, "关注失败: "+err.Error())
		}
	}

	controllers.Redirect(c, named.G("users.show", id)+"?page=1", false)
}

func Destroy(c *gin.Context, currentUser *userModel.User) {
	id, err := controllers.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	isFollowing := true
	if id != int(currentUser.ID) {
		isFollowing = followerModel.IsFollowing(int(currentUser.ID), id)
	}

	if isFollowing{
		if err := followerModel.DoUnFollow(currentUser.ID, uint(id)); err != nil {
			flash.NewDangerFlash(c, "取消关注失败: "+err.Error())
		}
	}
}
