package router

import (
	"core-sdk-example/api"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	sdkApi := router.Group("/api")
	// AddNode 添加节点
	sdkApi.POST("/addNode", api.AddNode)
	// RemoveNode 删除节点
	sdkApi.POST("/removeNode", api.RemoveNode)
	// RemoveAccount 删除账户
	sdkApi.POST("/removeAccount", api.RemoveAccount)
	// GetNodeState 查询节点状态
	sdkApi.GET("/getNodeState", api.GetNodeState)
	// GetNodeServerState 查询服务器状态
	sdkApi.GET("/getNodeServerState", api.GetNodeServerState)
	// GetNodeServerInfo 查询服务器信息
	sdkApi.GET("/getNodeServerInfo", api.GetNodeServerInfo)
}
