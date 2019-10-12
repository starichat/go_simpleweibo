package blog

import (
	"fmt"
	"go_simpleweibo/database"
	"strconv"
)

func Get(id int) (*Blog, error) {
	b := &Blog{}
	d := database.DB.First(&b, id)
	return b, d.Error
}

// 获取指定用户（们）的微博数量
func GetBlogsCountByUsers(ids []uint) (int, error) {
	sqlStr := fmt.Sprintf("select count(*) from %s where deleted_at is null and user_id in (", tableName)
	l := len(ids) - 1
	for i, v := range ids {
		sqlStr = sqlStr + strconv.Itoa(int(v))
		if i < l {
			sqlStr = sqlStr + ","
		}
	}
	sqlStr = sqlStr + ")"

	count := 0
	d := database.DB.Raw(sqlStr).Count(&count)
	return count, d.Error
}

// 获取指定用户（们）的微博
func GetBlogsByUsers(ids []uint, offset, limit int) ([]*Blog, error) {
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

// 根据blog_id 获取该微博的所有者
func GetUserByBlogID(blogID int) (*userModel.User, error) {
	b, err := Get(blogID)
	if err != nil {
		// 错误解析
		return nil, err
	}

	u, err := userModel.Get(int(b.UserId))
	if err != nil {
		return nil, err
	}

	return u, nil
}

// 获取该用户的微博，分页
func GetAllBlogsByUser(userID, offset, limit int) ([]*Blog, error) {
	blog := make([]*Blog, 0)
	err := database.DB.Where("user_id = ?", userID).Offset(offset).Limit(limit).Order("desc").Find(&blog).Error

	if err != nil {
		return blog, err
	}

	return blog, nil
}

// 获取该用户的所有微博数量
func GetBlogCountByUser(userID int) (count int, err error) {
	err = database.DB.Model(&Blog{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}
