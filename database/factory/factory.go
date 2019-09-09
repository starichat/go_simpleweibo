package factory

import "go_simpleweibo/database"

//　Drop　And Creat table
func DropAndCreateTable(table interface{}) {
	database.DB.DropTable(table)
	database.DB.CreateTable(table)
}
