package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 将文件名 MD5 加密后再进行写入，防止直接把原始名称暴露出去
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
