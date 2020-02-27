package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
