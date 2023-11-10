package api

import (
	"github.com/AiLiaa/blog-service/global"
	"github.com/AiLiaa/blog-service/internal/service"
	"github.com/AiLiaa/blog-service/pkg/app"
	"github.com/AiLiaa/blog-service/pkg/convert"
	"github.com/AiLiaa/blog-service/pkg/errcode"
	"github.com/AiLiaa/blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func NewUploadController() UploadController {
	return UploadController{}
}

func (u UploadController) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)

	//读取入参 file 字段的上传文件信息
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
