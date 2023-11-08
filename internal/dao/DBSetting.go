package dao

import (
	"fmt"
	"github.com/AiLiaa/blog-service/global"
	"github.com/AiLiaa/blog-service/utils/setting"
	"github.com/jinzhu/gorm"
)

// NewDBEngine 创建 DB 实例
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	//s1 := "root:root@tcp(127.0.0.1:3306)/blog_service?charset=utf8&parseTime=True&loc=Local"
	//db, err := gorm.Open("mysql", s1)

	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"

	//fmt.Println(fmt.Sprintf(s,
	//	databaseSetting.UserName,
	//	databaseSetting.Password,
	//	databaseSetting.Host,
	//	databaseSetting.DBName,
	//	databaseSetting.Charset,
	//	databaseSetting.ParseTime))

	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, err
}
