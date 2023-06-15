package network_utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-network-watcher/internal/global/g_consts"
	"time"
)

type uProxyNetwork struct{}

var ProxyNetwork = &uProxyNetwork{}

// GetProxyNetwork
//
//	@dc: 获取代理服务器的网速
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/2 20:06:21
func (u *uProxyNetwork) GetProxyNetwork() (err error) {
	var (
		proxyNetwork = g.Map{
			"time":        "",
			"rxSpeedKbps": 0,
			"txSpeedKbps": 0,
			"rxSpeedMbps": 0,
			"txSpeedMbps": 0,
		}
		url = g_consts.XuiBaseUrl + "/server/status"
	)

	session, err := u.GetXuiSession()
	// 通过xui进行网速的获取

	post, err := g.Client().SetCookieMap(session).Post(context.Background(), url)
	defer func(post *gclient.Response) {
		err := post.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(post)
	if err != nil {
		return err
	}
	if post.StatusCode != 200 {
		glog.Warning(context.Background(), "获取网速失败")
		return err
	}
	jsonData := gjson.New(post.ReadAllString())
	rxSpeed := jsonData.Get("obj.netIO.down") // 下载速度
	txSpeed := jsonData.Get("obj.netIO.up")   // 上传速度

	// 速度单位转换
	rxSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024))
	txSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024))
	proxyNetwork["rxSpeedKbps"] = rxSpeedKbps
	proxyNetwork["txSpeedKbps"] = txSpeedKbps

	// 转换成MB
	rxSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024/1024))
	txSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024/1024))
	proxyNetwork["rxSpeedMbps"] = rxSpeedMbps
	proxyNetwork["txSpeedMbps"] = txSpeedMbps

	proxyNetwork["time"] = gtime.Now().String()
	err = gcache.Set(context.Background(), g_consts.ProxyNetworkCacheKey, proxyNetwork, 0)
	if err != nil {
		return err
	}

	return err
}

// GetXuiSession
//
//	@dc: 获取Xui登陆session
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/2 20:06:21
func (u *uProxyNetwork) GetXuiSession() (sessionMap map[string]string, err error) {
	var (
		url = g_consts.XuiBaseUrl + "/login"
	)
	post, err := g.Client().Post(context.Background(), url, g_consts.XuiLoginDataMap)
	defer func(post *gclient.Response) {
		err := post.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(post)
	if err != nil {
		return nil, err
	}
	if post.StatusCode != 200 {
		return nil, fmt.Errorf("登录失败")
	}
	if post.Header.Get("Set-Cookie") == "" {
		return nil, fmt.Errorf("获取Cookie失败")
	}
	// 将session存入缓存
	err = gcache.Set(context.Background(), g_consts.ProxySessionCacheKey, post.GetCookieMap(), 15*time.Minute)
	if err != nil {
		return nil, err
	}
	return post.GetCookieMap(), nil
}
