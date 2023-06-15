package controller

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"net/http"
	"push/utility/network_utils"
	"testing"
	"time"
)

// 获取家庭路由器网速
func Test16_59_12(t *testing.T) {
	url := "http://120.24.211.49:35600/json/stats.json"
	response, err := g.Client().Get(context.Background(), url)
	if err != nil {
		return
	}
	jsonData := gjson.New(response.ReadAllString())
	rxSpeed := jsonData.Get("servers.0.network_rx") // 下载速度
	txSpeed := jsonData.Get("servers.0.network_tx") // 上传速度
	// 速度单位转换
	rxSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024))
	txSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024))
	t.Log("下载速度：", rxSpeedKbps, "KB/S")
	t.Log("上传速度：", txSpeedKbps, "KB/S")
	// 转换成MB
	rxSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024/1024))
	txSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024/1024))
	t.Log("下载速度：", rxSpeedMbps, "MB/S")
	t.Log("上传速度：", txSpeedMbps, "MB/S")
}

// 获取科学上网网速
func Test16_59_13(t *testing.T) {
	url := "http://120.24.211.49:35600/json/stats.json"
	response, err := g.Client().Get(context.Background(), url)
	if err != nil {
		return
	}
	jsonData := gjson.New(response.ReadAllString())
	rxSpeed := jsonData.Get("servers.1.network_rx") // 下载速度
	txSpeed := jsonData.Get("servers.1.network_tx") // 上传速度
	// 速度单位转换
	rxSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024))
	txSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024))
	t.Log("下载速度：", rxSpeedKbps, "KB/S")
	t.Log("上传速度：", txSpeedKbps, "KB/S")
	// 转换成MB
	rxSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024/1024))
	txSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024/1024))
	t.Log("下载速度：", rxSpeedMbps, "MB/S")
	t.Log("上传速度：", txSpeedMbps, "MB/S")
}

// 测试方法注释
func Test18_29_04(t *testing.T) {
	// 登陆获取token
	url := "http://ray.xinyu.today:580/api/login"
	jsonData := `{"username":"hamster","password":"deny1963"}`
	response, err := g.Client().Post(context.Background(), url, jsonData)
	if err != nil {
		return
	}
	resData := gjson.New(response.ReadAllString())
	if resData.Get("code").String() != "SUCCESS" {
		return
	}
	token := resData.Get("data.token").String()
	g.Dump(token)
	// websocket获取节点列表
	wsUrl := "ws://ray.xinyu.today:580/api/message?Authorization=" + token
	client := gclient.NewWebSocket()
	client.HandshakeTimeout = time.Second    // 设置超时时间
	client.Proxy = http.ProxyFromEnvironment // 设置代理
	client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置
	conn, _, err := client.Dial(wsUrl, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var connectedNode string
	var lastChangeTime string
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			panic(err)
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
					g.Dump(gjson.New(connectedNode).Get("delay").Int())
					g.Dump(gjson.New(gconv.String(v)).Get("delay").Int())
					connectedNode = gconv.String(v)
					lastChangeTime = gtime.Now().String()
				}
			}
		}
		g.Dump(gjson.New(connectedNode).Get("outbound_tag").String())
		g.Dump(lastChangeTime)
	}

}

// 测试方法注释
func Test20_09_22(t *testing.T) {
	count := 0
	homeNetwork := g.Map{
		"time":        "",
		"rxSpeedKbps": 0,
		"txSpeedKbps": 0,
		"rxSpeedMbps": 0,
		"txSpeedMbps": 0,
	}
	g.Dump("开始获取家庭路由器网速")
	for {
		if count%2 == 0 {
			url := "http://120.24.211.49:35600/json/stats.json"
			response, err := g.Client().Get(context.Background(), url)
			if err != nil {
				continue
			}
			jsonData := gjson.New(response.ReadAllString())
			rxSpeed := jsonData.Get("servers.0.network_rx") // 下载速度
			txSpeed := jsonData.Get("servers.0.network_tx") // 上传速度
			// 速度单位转换
			rxSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024))
			txSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024))
			homeNetwork["rxSpeedKbps"] = rxSpeedKbps
			homeNetwork["txSpeedKbps"] = txSpeedKbps
			// 转换成MB
			rxSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024/1024))
			txSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024/1024))
			homeNetwork["rxSpeedMbps"] = rxSpeedMbps
			homeNetwork["txSpeedMbps"] = txSpeedMbps
			homeNetwork["time"] = gtime.Now().String()
			err = gcache.Set(context.Background(), "homeNetwork", homeNetwork, 0)
			if err != nil {
				return
			}
			g.Dump(homeNetwork)
			count += 2
			time.Sleep(time.Second * 1)
		}
	}
}

// 测试方法注释
func Test20_20_30(t *testing.T) {
	count := 0
	token, err := network_utils.NodeUtils.GetToken()
	if err != nil {
		return
	}
	g.Dump(token)
	var connectedNode string
	var lastChangeTime string
	for {
		if count%50 == 0 {
			token, err = network_utils.NodeUtils.GetToken()
			if err != nil {
				return
			}
			g.Dump(token)
		}
		if count%10 == 0 {
			// websocket获取节点列表
			wsUrl := "ws://ray.xinyu.today:580/api/message?Authorization=" + token
			client := gclient.NewWebSocket()
			client.HandshakeTimeout = time.Second    // 设置超时时间
			client.Proxy = http.ProxyFromEnvironment // 设置代理
			client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置
			conn, _, err := client.Dial(wsUrl, nil)
			if err != nil {
				panic(err)
			}
			defer conn.Close()
			_, data, err := conn.ReadMessage()
			if err != nil {
				panic(err)
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
						g.Dump(gjson.New(connectedNode).Get("delay").Int())
						g.Dump(gjson.New(gconv.String(v)).Get("delay").Int())
						connectedNode = gconv.String(v)
						lastChangeTime = gtime.Now().String()
					}
				}
			}
			g.Dump(gjson.New(connectedNode).Get("outbound_tag").String())
			g.Dump(lastChangeTime)
		}

		count += 1
		time.Sleep(time.Second * 1)
	}
}
