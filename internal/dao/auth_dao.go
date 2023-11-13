package dao

import (
	"github.com/AiLiaa/blog-service/internal/model"
	"github.com/jinzhu/gorm"
)

func (dao *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	var auth model.Auth
	db := dao.engine

	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", appKey, appSecret, 0)
	if err := db.First(&auth).Error; err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
