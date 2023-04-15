package api

import (
	"core-sdk-example/dao/redis"
	"core-sdk-example/module/vo"
	"core-sdk-example/util"
	"coresdk"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// 在调用API之前需要在Redis中设置trojan-panel:jwk-key和trojan-panel:token:${username}
func prepareRequest() (string, error) {
	username := "sysadmin"
	tokenStr, err := util.GenToken(vo.AccountVo{})
	if err != nil {
		return "", err
	} else {
		if _, err = redis.Client.String.
			Set(fmt.Sprintf("trojan-panel:token:%s", username), tokenStr,
				time.Hour.Milliseconds()*2/1000).Result(); err != nil {
			return "", err
		}
		return tokenStr, nil
	}
}

// AddNode 添加节点
func AddNode(c *gin.Context) {
	token, err := prepareRequest()
	if err != nil {
		logrus.Errorf("prepareRequest err: %v", err)
		return
	}
	dto := coresdk.NodeAddDto{
		NodeTypeId: 4,
		Port:       443,
		Domain:     "example.com",
	}
	if err := coresdk.AddNode(token,
		"127.0.0.1", 8100, &dto); err != nil {
		logrus.Errorf("AddNode err: %v", err)
		return
	}
	logrus.Errorf("AddNode success")
}

// RemoveNode 删除节点
func RemoveNode(c *gin.Context) {
	token, err := prepareRequest()
	if err != nil {
		logrus.Errorf("prepareRequest err: %v", err)
		return
	}
	removeDto := coresdk.NodeRemoveDto{NodeTypeId: 4, Port: 443}
	if err := coresdk.RemoveNode(token,
		"127.0.0.1", 8100, &removeDto); err != nil {
		logrus.Errorf("RemoveNode err: %v", err)
		return
	}
	logrus.Errorf("RemoveNode success")
}

// RemoveAccount 删除账户
func RemoveAccount(c *gin.Context) {
	token, err := prepareRequest()
	if err != nil {
		logrus.Errorf("prepareRequest err: %v", err)
		return
	}
	accountRemoveDto := coresdk.AccountRemoveDto{}
	if err := coresdk.RemoveAccount(token,
		"127.0.0.1", 8100, &accountRemoveDto); err != nil {
		logrus.Errorf("RemoveAccount err: %v", err)
		return
	}
	logrus.Errorf("RemoveAccount success")
}

// GetNodeState 查询节点状态
func GetNodeState(c *gin.Context) {
	token, err := prepareRequest()
	if err != nil {
		logrus.Errorf("prepareRequest err: %v", err)
		return
	}
	nodeState, err := coresdk.GetNodeState(token,
		"127.0.0.1", 8100, 4, 443)
	if err != nil {
		logrus.Errorf("GetNodeState err: %v", err)
		return
	}
	logrus.Errorf("GetNodeState success result: %v", nodeState)
}

// GetNodeServerState 查询服务器状态
func GetNodeServerState(c *gin.Context) {
	token, err := prepareRequest()
	if err != nil {
		logrus.Errorf("prepareRequest err: %v", err)
		return
	}
	nodeServerState, err := coresdk.GetNodeServerState(token,
		"127.0.0.1", 8100)
	if err != nil {
		logrus.Errorf("GetNodeServerState err: %v", err)
		return
	}
	logrus.Errorf("GetNodeServerState success result: %v", nodeServerState)
}

// GetNodeServerInfo 查询服务器信息
func GetNodeServerInfo(c *gin.Context) {
	token, err := prepareRequest()
	if err != nil {
		logrus.Errorf("prepareRequest err: %v", err)
		return
	}
	nodeServerInfo, err := coresdk.GetNodeServerInfo(token,
		"127.0.0.1", 8100)
	if err != nil {
		logrus.Errorf("GetNodeServerInfo err: %v", err)
		return
	}
	logrus.Errorf("GetNodeServerInfo success result: %v", nodeServerInfo)
}
