package models

import (
	"fmt"

	"../libs/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	CommonModel struct {
		gorm.Model
	}
)

var (
	DB *gorm.DB
)

func Init() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		config.Get().Db.User,
		config.Get().Db.Password,
		config.Get().Db.Dbname))

	if err != nil {
		panic("数据库连接错误")
	}

}
