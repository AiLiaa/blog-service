package global

import (
	"github.com/AiLiaa/blog-service/pkg/logger"
	"github.com/AiLiaa/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

// 全局变量 将配置信息和应用程序关联起来
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
)
