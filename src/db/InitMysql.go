package db

import (
	"bk/src/config"
	"bk/src/static"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql(config *config.MysqlConfig) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", config.UserName, config.Password, config.Addr, config.Db)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection failure")
	}
	DB = db

	//设置默认数据库引擎
	_ = DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		//初始化用户数据表
		&static.UserInfos{},
		&static.TestQuestions{},
		&static.Log{},
		)
}
