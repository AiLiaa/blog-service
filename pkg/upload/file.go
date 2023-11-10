package upload

import (
	"github.com/AiLiaa/blog-service/global"
	"github.com/AiLiaa/blog-service/pkg/util"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

// GetFileName 通过获取文件后缀并筛出原始文件名进行 MD5 加密，最后返回经过加密处理后的文件名。
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetFileExt 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取文件保存地址
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath 检查保存目录是否存在
// 利用 os.Stat 方法所返回的 error 值与系统中所定义的 oserror.ErrNotExist 进行判断
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckContainExt 检查文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

// CheckMaxSize 检查文件大小是否超出最大大小限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

// CheckPermission 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// CreateSavePath 创建在上传文件时所使用的保存目录
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// SaveFile 保存所上传的文件
// 调用 os.Create 方法创建目标地址的文件，再通过 file.Open 方法打开源地址的文件，结合 io.Copy 方法实现两者之间的文件内容拷贝。
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
