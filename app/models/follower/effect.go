package follower

import (
	"fmt"
	"go_simpleweibo/database"
	"strconv"
)


// 执行关注
func DoFollow(userID uint, followIDs ...uint) error {
	fmt.Println(followIDs)
	// 获取被关注者列表的长度
	l := len(followIDs) - 1
	fmt.Println(l)
	// 构造执行关注的insert  sql语句
	sqlStr := fmt.Sprintf("insert into %s (follower_id, user_id) values ", tableName)
	// 遍历关注所有偶像
	for i, v := range followIDs {
		fmt.Println("aaaaaaaa")
		sqlStr = sqlStr + fmt.Sprintf("(%d, %d)", userID, v)
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	fmt.Println(sqlStr)
	d := database.DB.Exec(sqlStr)
	return d.Error
}

// 执行取消关注
func DoUnFollow(userID uint, followIDs ...uint) error {
	sqlStr := fmt.Sprintf("delete from %s where follower_id = %d and user_id in (", tableName, userID)
	l := len(followIDs) - 1
	for i, v := range followIDs {
		sqlStr = sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + ")"
	fmt.Println(sqlStr)
	d := database.DB.Exec(sqlStr)
	return d.Error
}