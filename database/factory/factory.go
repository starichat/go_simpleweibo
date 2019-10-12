package factory

import "go_simpleweibo/database"

//　Drop　And Creat table
func DropAndCreateTable(table interface{}) {
	// database.DB 为database中的gorm类型的DB数据库
	// gorm 提供了 CreateTable（），Find（），Where（），DropTable（）
	database.DB.DropTable(table)
	database.DB.CreateTable(table)
}
