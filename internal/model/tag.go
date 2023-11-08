package model

import "github.com/AiLiaa/blog-service/utils/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// TagSwagger 用于 Swagger 接口文档展示
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
