package network_utils

import (
	"context"
	"crypto/tls"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-network-watcher/internal/global/g_consts"
	"net/http"
	"time"
)

type uNodeUtils struct{}

var NodeUtils = &uNodeUtils{}

// GetNodeInfo
//
//	@dc:
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/2 20:08:50
func (u *uNodeUtils) GetNodeInfo() (err error) {
	var (
		connectedNode  string
		lastChangeTime string
		nodeInfo       = g.Map{
			"nodeName":       "",
			"lastChangeTime": "",
			"updateTime":     "",
		}
	)

	// 获取token
	token, err := u.GetToken()
	if err != nil {
		return err
	}
	wsUrl := "ws://" + g_consts.XrayBaseUrl + "/api/message?Authorization=" + token

	// websocket获取节点列表
	client := gclient.NewWebSocket()
	client.HandshakeTimeout = time.Second    // 设置超时时间
	client.Proxy = http.ProxyFromEnvironment // 设置代理
	client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置
	conn, _, err := client.Dial(wsUrl, nil)
	if err != nil {
		return err
	}
	_, data, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	// 打印消息类型和消息内容
	nodeList := gjson.New(string(data)).Get("body.outboundStatus").Array()
	for _, v := range nodeList {
		if !gjson.New(gconv.String(v)).Get("alive").Bool() {
			continue
		}
		// 根据"delay"获得延迟最小的节点
		if connectedNode == "" {
			connectedNode = gconv.String(v)
		} else {
			if gjson.New(connectedNode).Get("delay").Int() > gjson.New(gconv.String(v)).Get("delay").Int() {
				connectedNode = gconv.String(v)
			}
		}
	}
	// 比对缓存中的节点信息
	cacheNodeInfo, _ := gcache.Get(context.Background(), g_consts.ProxyNodeCacheKey)
	if cacheNodeInfo != nil {
		if gjson.New(cacheNodeInfo.String()).Get("nodeName").String() == gjson.New(connectedNode).Get("outbound_tag").String() {
			lastChangeTime = gjson.New(cacheNodeInfo).Get("lastChangeTime").String()
		} else {
			lastChangeTime = gtime.Now().String()
		}
	} else {
		lastChangeTime = gtime.Now().String()
	}
	// 获取节点信息
	nodeInfo["nodeName"] = gjson.New(connectedNode).Get("outbound_tag").String()
	nodeInfo["lastChangeTime"] = lastChangeTime
	nodeInfo["updateTime"] = gtime.Now().String()
	// 关闭连接
	_ = conn.Close()
	// 存入缓存
	err = gcache.Set(context.Background(), g_consts.ProxyNodeCacheKey, nodeInfo, 0)
	if err != nil {
		return err
	}
	return nil
}

// GetToken
//
//	@dc: 获取token
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/2 20:13:24
func (u *uNodeUtils) GetToken() (token string, err error) {
	var (
		url = "http://" + g_consts.XrayBaseUrl + "/api/login"
	)

	// 登陆获取token
	response, err := g.Client().Post(context.Background(), url, g_consts.XrayLoginDataMap)
	if err != nil {
		return "", err
	}
	resData := gjson.New(response.ReadAllString())
	if resData.Get("code").String() != "SUCCESS" {
		err = gerror.New("登陆失败")
		return
	}
	token = resData.Get("data.token").String()
	if token == "" {
		err = gerror.New("token获取失败")
		return
	}
	return
}
