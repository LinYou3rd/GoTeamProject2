package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbUser *gorm.DB

func Init() {

	var err error
	var constr string
	constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "tianxiaDIYI427", "localhost", 3306, "raise_money")
	dbUser, err = gorm.Open("mysql", constr)
	if err != nil {
		println(err)
		panic("存放用户数据的mysql数据库连接失败")
	}
	dbUser.AutoMigrate(&userModel{})
	dbUser.AutoMigrate(&adminModel{})

}
