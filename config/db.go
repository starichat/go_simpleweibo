package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// 数据库配置
type dbConfig struct {
	Connection string // 连接
	Host       string //主机
	Port       int    // 端口
	Database   string // 数据库名
	Username   string // 用户名
	Password   string // 密码

	URL string // URL

	Debug bool // 是否debug
}

func newDBConfig() *dbConfig {
	// 默认配置
	viper.SetDefault("DB.CONNECTION", "mysql")
	viper.SetDefault("DB.HOST", "localhost")
	viper.SetDefault("DB.PORT", 3306)
	viper.SetDefault("DB.DATABASE", "go_simpleweibo")
	viper.SetDefault("DB.USERNAME", "root")
	viper.SetDefault("DB.PASSWORD", "")

	username := viper.GetString("DB.USERNAME")
	password := viper.GetString("DB.PASSWORD")
	host := viper.GetString("DB.HOST")
	port := viper.GetInt("DB.PORT")
	database := viper.GetString("DB.DATABASE")
	url := createDBURL(username, password, database)

	return &dbConfig{
		Connection: viper.GetString("DB.CONNECTION"),
		Host:       host,
		Port:       port,
		Database:   database,
		Username:   username,
		Password:   password,
		URL:        url,
		Debug:      AppConfig.RunMode == RunmodeDebug,
	}

}

func createDBURL(uname string, pwd string, name string) string {
	return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=%t&loc=%s",
		uname, pwd,
		name, true, "Local")
}
