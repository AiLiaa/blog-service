package v1

import (
	"github.com/AiLiaa/blog-service/pkg/app"
	"github.com/AiLiaa/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func NewArticle() ArticleController {
	return ArticleController{}
}

// @Summary 获取单个文章
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1//articles/{id} [get]
func (a ArticleController) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 获取多个文章
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1//articles [get]
func (a ArticleController) List(c *gin.Context) {}

// @Summary 创建文章
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1//articles [post]
func (a ArticleController) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1//articles/{id} [patch]
func (a ArticleController) Update(c *gin.Context) {}

// @Summary 更新文章
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1//articles/{id} [delete]
func (a ArticleController) Delete(c *gin.Context) {}
