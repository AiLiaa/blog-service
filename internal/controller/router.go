package controller

import (
	_ "github.com/AiLiaa/blog-service/docs"
	"github.com/AiLiaa/blog-service/global"
	"github.com/AiLiaa/blog-service/internal/controller/api"
	v1 "github.com/AiLiaa/blog-service/internal/controller/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   //gin日志
	r.Use(gin.Recovery()) //gin恢复
	//r.Use(config.Translations()) //validator 参数校验翻译中间件（这样会重复注册）

	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	articleController := v1.NewArticleController()
	tagController := v1.NewTagController()
	uploadController := api.NewUploadController()

	//上传文件
	r.POST("/upload/file", uploadController.UploadFile)
	//访问静态资源
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tagController.Create)
		apiv1.DELETE("/tags/:id", tagController.Delete)
		apiv1.PUT("/tags/:id", tagController.Update)
		apiv1.PATCH("/tags/:id/state", tagController.Update)
		apiv1.GET("/tags", tagController.List)

		apiv1.POST("/articles", articleController.Create)
		apiv1.DELETE("/articles/:id", articleController.Delete)
		apiv1.PUT("/articles/:id", articleController.Update)
		apiv1.PATCH("/articles/:id/state", articleController.Update)
		apiv1.GET("/article", articleController.Get)
		apiv1.GET("/articles", articleController.List)
	}

	return r
}
