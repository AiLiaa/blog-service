package model

import "github.com/AiLiaa/blog-service/pkg/app"

type Article struct {
	*Common
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

// ArticleSwagger 用于 Swagger 接口文档展示
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
