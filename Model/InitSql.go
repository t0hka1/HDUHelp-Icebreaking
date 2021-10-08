package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID uint
	StudentId string
	Password string
	Name string
	PicPath string
	Department string
}
var DB *gorm.DB

func Mysql1()(err error)  {
	DB,_=gorm.Open("mysql", "root:root@/ice_breaking?charset=utf8&parseTime=True&loc=Local")
	DB.AutoMigrate(&User{})
	return err
}
