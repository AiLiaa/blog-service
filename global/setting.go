package global

import (
	"github.com/AiLiaa/blog-service/utils/logger"
	"github.com/AiLiaa/blog-service/utils/setting"
	"github.com/jinzhu/gorm"
)

// 全局变量 将配置信息和应用程序关联起来
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
