package service

import (
	"github.com/AiLiaa/blog-service/internal/model"
	"github.com/AiLiaa/blog-service/pkg/app"
)

type CreateArticleRequest struct {
	Title     string `form:"title" binding:"required,min=1,max=100"`
	Desc      string `form:"desc" binding:"required,min=1,max=200"`
	Content   string `form:"content" binding:"required,min=1,max=1024"`
	CreatedBy string `form:"created_by" binding:"required"`
}

type UpdateArticleRequest struct {
	Id         uint32 `form:"id" binding:"required,gte=1"`
	Title      string `form:"title" binding:"required,min=1,max=100"`
	Desc       string `form:"desc" binding:"required,min=1,max=200"`
	Content    string `form:"content" binding:"required,min=1,max=1024"`
	ModifiedBy string `form:"modified_by" binding:"required"`
}

type DeleteArticleRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CreatedBy)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.Id, param.Title, param.Desc, param.Content, param.ModifiedBy)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.Id)
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return svc.dao.GetArticleCount(param.Title, param.State)
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) GetArticle(param *GetArticleRequest) (model.Article, error) {
	return svc.dao.GetArticle(param.Title, param.State)
}
