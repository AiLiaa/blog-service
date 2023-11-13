package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimiterIface interface {
	// Key 获取对应的限流器的键值对名称
	Key(c *gin.Context) string
	// GetBucket 获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	// AddBuckets 新增多个令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

// Limiter 存储令牌桶与键值对名称的映射关系
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// LimiterBucketRule 存储令牌桶的一些相应规则属性
type LimiterBucketRule struct {
	Key          string        //自定义键值对名称
	FillInterval time.Duration //间隔多久时间放 N 个令牌
	Capacity     int64         //令牌桶的容量
	Quantum      int64         //每次到达间隔时间后所放的具体令牌数量
}
