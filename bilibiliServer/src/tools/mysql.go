package tools

import (
	"bilibiliServer/src/model"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() error {
	// 用户名:密码@tcp(ip:port)/数据库名字?charset=utf8&parseTime=True
	dsn := Config.DbUsername + ":" + Config.DbPassword + "@tcp(127.0.0.1:3306)/" + Config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("数据库打开失败")
	}
	DB = Db

	Db.AutoMigrate(&model.User{})
	Db.AutoMigrate(&model.Video{})
	Db.AutoMigrate(&model.UserFollow{})
	Db.AutoMigrate(&model.VideoLike{})
	Db.AutoMigrate(&model.Sms_record{})
	DB.AutoMigrate(&model.Comment{})
	return nil
}

func GetDb() *gorm.DB {
	return DB
}
