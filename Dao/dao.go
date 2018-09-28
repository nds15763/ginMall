package dao

import (
	"ginMall/Model"

	"github.com/jinzhu/gorm"
)

type DaoManager struct {
	DB *gorm.DB
	//logger *logger.Logger
}

func NewDaoManager(constr string) (*DaoManager, error) {

	db, err := gorm.Open("mysql", constr)
	if err != nil {
		return nil, err
	}

	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// ORM自动迁移模式
	db.AutoMigrate(&Model.UserModel{},
		&Model.UserDetailModel{},
		&Model.UserAuthsModel{},
	)

	dao := &DaoManager{
		DB: db,
		//logger: logger,
	}
	return dao, nil
}
func (this *DaoManager) GetEngine() *gorm.DB {
	return this.DB
}
