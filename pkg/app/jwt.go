package app

import (
	"github.com/AiLiaa/blog-service/global"
	"github.com/AiLiaa/blog-service/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	//	Audience  string `json:"aud,omitempty"`
	//	ExpiresAt int64  `json:"exp,omitempty"`
	//	Id        string `json:"jti,omitempty"`
	//	IssuedAt  int64  `json:"iat,omitempty"`
	//	Issuer    string `json:"iss,omitempty"`
	//	NotBefore int64  `json:"nbf,omitempty"`
	//	Subject   string `json:"sub,omitempty"`
	jwt.StandardClaims
}

// GetJWTSecret 获取该项目的 JWT Secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成 JWT Token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	//根据 Claims 结构体创建 Token 实例
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成签名字符串
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// ParseToken 解析和校验 Token
func ParseToken(token string) (*Claims, error) {
	//解析鉴权的声明 返回 *Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//Valid 验证基于时间的声明
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
