package service

import (
	"ginMall/logger"
	"time"

	"ginMall/config"
	"ginMall/dao"
	"ginMall/httpsvr"
)

type Service struct {
	cfgFlag string // 配置文件路径

	config   *config.Configuration //系统配置
	logger   *logger.Logger        //日志
	location *time.Location        // 时区信息
	httpsvr  *httpsvr.HttpServer

	daomgr *dao.DaoManager //引用Dao层

	//redismgr *redis.RedisManager //Redis
}

func NewService(params *config.Configuration, logger *logger.Logger, location *time.Location, cfg string) *Service {

	// 初始化数据库连接
	dsn := params.DbUser + ":" + params.DbPasswd + "@tcp" + "(" + params.DbUrl + ":" + params.DbPort + ")" + "/" + params.DbName + "?charset=utf8"

	// 初始化Dao层
	daomgr, err := dao.NewDaoManager(dsn)
	if err != nil {
		logger.Error(err)
	}

	httpsvr := httpsvr.NewHttpServer()
	// if err != nil {
	// 	logger.Error(err)
	// }

	service := &Service{

		cfgFlag:  cfg,
		config:   params,
		logger:   logger,
		location: location,
		daomgr:   daomgr,
		httpsvr:  httpsvr,
	}
	return service
}
func (this *Service) Start() {
	go this.httpsvr.Start()
}
