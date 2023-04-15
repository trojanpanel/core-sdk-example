package main

import (
	"core-sdk-example/core"
	"core-sdk-example/dao/redis"
	"core-sdk-example/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Router(r)
	_ = r.Run(":8080")
	defer closeResource()
}

func init() {
	core.InitConfig()
	redis.InitRedis()
}

func closeResource() {
	redis.CloseRedis()
}
