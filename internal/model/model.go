package model

import (
	"fmt"
	"github.com/AiLiaa/blog-service/global"
	"github.com/AiLiaa/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	//引入 MySQL 驱动库
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

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
