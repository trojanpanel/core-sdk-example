package util

import (
	"core-sdk-example/dao/redis"
	"core-sdk-example/module/constant"
	"core-sdk-example/module/vo"
	"errors"
	"github.com/golang-jwt/jwt"
	redisgo "github.com/gomodule/redigo/redis"
	"time"
)

// TokenExpireDuration 过期时间默认2小时
const TokenExpireDuration = time.Hour * 2

type MyClaims struct {
	AccountVo vo.AccountVo `json:"accountVo"`
	jwt.StandardClaims
}

// GenToken 生成Token
func GenToken(accountVo vo.AccountVo) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		// 自定义字段
		AccountVo: accountVo,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			// 签发人
			Issuer: constant.SystemName,
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	mySecret, err := GetJWTKey()
	if err != nil {
		return "", errors.New(constant.SysError)
	}
	return token.SignedString(mySecret)
}

func GetJWTKey() ([]byte, error) {
	get := redis.Client.String.
		Get(constant.JwtKey)
	reply, err := get.Bytes()
	if err != nil && err != redisgo.ErrNil {
		return nil, errors.New(constant.SysError)
	}
	if len(reply) > 0 {
		return reply, nil
	} else {
		// jwt key 72小时更新一次
		key := []byte(RandString(32))
		_, err := redis.Client.String.Set(constant.JwtKey, key, time.Hour.Milliseconds()*72/1000).Result()
		if err != nil {
			return nil, errors.New(constant.SysError)
		}
		return key, nil
	}
}
