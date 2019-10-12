package follower

import (
	"fmt"
	"go_simpleweibo/database"
	"strconv"
)

const (
	tableName = "followers"
)

// Follower 粉丝
type Follower struct {
	ID         uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	UserID     uint `gorm:"column:user_id;not null" sql:"index"`     // 多对多，关联 User Model (关注者)
	FollowerID uint `gorm:"column:follower_id;not null" sql:"index"` // 多对多，关联 User Model (粉丝)
}

// TableName 表名
func (Follower) TableName() string {
	return tableName
}

// 关注
func DoFollow(userId uint, followIds ...uint) error {
	l := len(followIds) - 1
	sqlStr := fmt.Sprintf("insert into %s (follower_id, user_id) values ", tableName)
	for i, v := range followIds {
		sqlStr = sqlStr + fmt.Sprintf("(%d, %d)", userId, v)
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	d := database.DB.Exec(sqlStr)
	return d.Error
}

// 取消关注
func DoUnFollow(userID uint, follfollowIds ...uint) error {
	sqlStr := fmt.Sprintf("delete from %s where follower_id = %d and user_id in (", tableName, userId)
	l := len(follfollowIds) - 1
	for i, v := range followIds {
		sqlStr = sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + ")"
	d := database.DB.Exec(sqlStr)
	return d.Error
}
