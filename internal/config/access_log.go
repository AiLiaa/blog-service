package config

import (
	"bytes"
	"github.com/AiLiaa/blog-service/global"
	"github.com/AiLiaa/blog-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// AccessLogWriter 双写，可以直接通过 AccessLogWriter 的 body 取到值
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		//初始化 AccessLogWriter，将其赋予给当前的 Writer 写入流
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
