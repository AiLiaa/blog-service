package config

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		//设置当前 context 的超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		//重新赋予给 gin.Context
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
