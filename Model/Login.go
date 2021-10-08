package Model

import (
	"database/sql"
	"errors"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PWD struct {
	Password string
}
//登录时的参数与数据库中的数据进行校验

func TryLogin(StudentId string,Password string) (err error){
	pwd:=new(PWD)
	err=DB.Table("users").Select("password").Where("student_id=?",StudentId).First(pwd).Error
	if err==sql.ErrNoRows{
		return errors.New(" Wrong StudentId!")
	}
	if err==nil{
		return err
	}
	if Password!=pwd.Password{
		return errors.New(" Wrong Password!")
	}
	return nil

}
