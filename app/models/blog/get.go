package blog

import (
	"fmt"
	userModel "go_simpleweibo/app/models/user"
	"go_simpleweibo/database"
	"strconv"
)

// blog用户模型匹配多种查找方式

func Get(id int) (*Blog, error) {
	b := &Blog{}
	d := database.DB.First(&b, id)
	return b, d.Error
}

// 获取指定用户的微博数量
func GetBYUsersBlogCount(ids []uint) (int,error) {
	sqlStr := fmt.Sprintf("select count(*) from %s where deleted_at is null and user_id in(",tableName)
	l := len(ids) - 1
	for i,v:=range ids {
		sqlStr := sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + ")"

	count := 0
	d := database.DB.Raw(sqlStr).Count(&count)
	return count, d.Error
}

// GetByUsersStatuses 获取指定用户们的微博
func GetByUsersStatuses(ids []uint, offset, limit int) ([]*Blog, error) {
	blog := make([]*Blog, 0)

	sqlStr := fmt.Sprintf("select * from %s where deleted_at is null and user_id in (", tableName)
	l := len(ids) - 1
	for i, v := range ids {
		sqlStr = sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + fmt.Sprintf(") order by `created_at` desc limit %d offset %d", limit, offset)

	d := database.DB.Raw(sqlStr).Scan(&blog)
	return blog, d.Error
}

// GetUser 通过 status_id 获取该微博的所有者
func GetUser(blogID int) (*userModel.User, error) {
	s, err := Get(blogID)
	if err != nil {
		return nil, err
	}

	u, err := userModel.Get(int(s.UserID))
	if err != nil {
		return nil, err
	}

	return u, nil
}


// GetUserAllStatus 获取该用户的所有微博
func GetUserAllStatus(userID int) ([]*Blog, error) {
	blog := make([]*Blog, 0)

	err := database.DB.Where("user_id = ?", userID).Order("id desc").Find(&blog).Error

	if err != nil {
		return blog, err
	}

	return blog, nil
}

// GetUserStatus 获取该用户的微博 (分页)
func GetUserStatus(userID, offset, limit int) ([]*Blog, error) {
	blog := make([]*Blog, 0)

	err := database.DB.Where("user_id = ?", userID).Offset(
		offset).Limit(limit).Order("id desc").Find(&blog).Error

	if err != nil {
		return blog, err
	}

	return blog, nil
}

// GetUserAllStatusCount 获取该用户的所有微博 的 count
func GetUserAllStatusCount(userID int) (count int, err error) {
	err = database.DB.Model(&Blog{}).Where("user_id = ?", userID).Count(&count).Error
	return
}