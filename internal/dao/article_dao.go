package dao

import (
	"github.com/AiLiaa/blog-service/internal/model"
	"github.com/AiLiaa/blog-service/pkg/app"
)

// CreateArticle 创建文章
func (dao *Dao) CreateArticle(title string, desc string, content string, createdBy string) error {
	db := dao.engine
	article := model.Article{
		Title:   title,
		Desc:    desc,
		Content: content,
		State:   1,
		Common:  &model.Common{CreatedBy: createdBy},
	}

	return db.Create(&article).Error
}

// UpdateArticle 更新文章
func (dao *Dao) UpdateArticle(id uint32, title string, desc string, content string, modifiedBy string) error {
	db := dao.engine
	article := model.Article{
		Title:   title,
		Desc:    desc,
		Content: content,
		Common:  &model.Common{ModifiedBy: modifiedBy, ID: id},
	}

	return db.Model(&article).Where("id = ? and is_del = ?", article.ID, 0).Update(&article).Error
}

// DeleteArticle 删除文章
func (dao *Dao) DeleteArticle(id uint32) error {
	db := dao.engine
	article := model.Article{Common: &model.Common{ID: id}}
	return db.Where("id = ? and is_del = ?", article.ID, 0).Delete(&article).Error
}

// GetArticleCount 获取文章数量
func (dao *Dao) GetArticleCount(title string, state uint8) (int, error) {
	var count int
	article := model.Article{Title: title, State: state}
	db := dao.engine

	if article.Title != "" {
		db = db.Where("title = ?", article.Title)
	}
	db = db.Where("state = ?", article.State)
	if err := db.Model(&article).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticleList 获取多个文章
func (dao *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	var articles []*model.Article
	var err error
	article := model.Article{Title: title, State: state}
	db := dao.engine

	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if article.Title != "" {
		db = db.Where("title = ?", article.Title)
	}
	db = db.Where("state = ?", article.State)
	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

// GetArticle 获取一个文章
func (dao *Dao) GetArticle(title string, state uint8) (model.Article, error) {
	var article_res model.Article
	article := model.Article{Title: title, State: state}
	db := dao.engine

	db = db.Where("state = ? and title = ?", article.State, article.Title)
	if err := db.Where("is_del = ?", 0).First(&article_res).Error; err != nil {
		return article_res, err
	}
	return article_res, nil
}
