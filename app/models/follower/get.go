package follower

import (
	"fmt"
	"go_simpleweibo/database"
)

func getFollowers(userID, limit int) (followers []*userModel.User, err error) {
	followers = make([]*userModel.User, 0)
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.follower_id", tableName)
	if limit == 0 {
		d := database.DB.Model(&userModel.User{}.Joins(joinSQL).Where("followers.user_id = ?", userID).Order("id").Find(&followers))
		return followers, d.Error
	} else {
		d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.user_id = ?", userID).Offset(offset).Limit(limit).Order("id").Find(&followers)
		return followers, d.Error
	}
}

func GetFollowings(userID, offset int) (followers []*userModel.User, err error) {
	followers = make([]*userModel.User, 0)
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.user_id", tableName)
	if limit == 0 {
		d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Order("id").Find(&followers)
		return followers, d.Error
	}

	d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Offset(offset).Limit(limit).Order("id").Find(&followers)
	return followers, d.Error

}

// FollowingsCount 关注数

// FollowersCount 粉丝数

// IsFollowing 已经关注了
