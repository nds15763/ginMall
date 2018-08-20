package Helper

import (
	"../../GinMall/Model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Dbinit() *gorm.DB {
	db := NewConn()
	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 自动迁移模式
	db.AutoMigrate(&Model.UserModel{},
		&Model.UserDetailModel{},
		&Model.UserAuthsModel{},
	)
	return db
}
func NewConn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1qaz!QAZ@/goweb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	return db
}
