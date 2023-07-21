package net_utils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"kes-cron/internal/global/g_consts"
	"kes-cron/internal/global/g_structs"
	"net/http"
	"time"
)

type uNetUtils struct{}

var NetUtils = &uNetUtils{}

// HttpsLatency
//
//	@dc: 测试https延迟
//	@author: Hamster   @date:2023-06-17 14:01:06
func (u *uNetUtils) HttpsLatency() (latency int, err error) {
	start := time.Now()
	resp, err := http.Get(g_consts.PingUrl)
	if err != nil {
		glog.Warning(context.Background(), "请求出错:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			glog.Warning(context.Background(), "关闭请求时发生错误:", err)
		}
	}(resp.Body)
	latency = gconv.Int(time.Since(start).Milliseconds())
	return latency, nil
}

// CoreLatency
//
//	@dc: 测试延迟核心服务
//	@author: Hamster   @date:2023-06-17 14:03:41
func (u *uNetUtils) CoreLatency(initData *g_structs.InitData) (err error) {
	latency, err := u.HttpsLatency()
	if err != nil {
		glog.Warning(context.Background(), "请求出错:", err)
		return
	}
	glog.Info(context.Background(), "HTTPS延迟:", latency)
	err = u.PushLatencyToServer(initData, latency)
	if err != nil {
		glog.Warning(context.Background(), "推送延迟到服务器时发生错误:", err)
		return err
	}
	return
}

// PushLatencyToServer
//
//	@dc: 推送延迟到服务器
//	@author: Hamster   @date:2023-06-17 14:02:59
func (u *uNetUtils) PushLatencyToServer(initData *g_structs.InitData, latency int) (err error) {
	_, macAddress := NetworkInfo.GetMacAddress()
	params := g.Map{
		"department":  initData.Department,
		"staff_name":  initData.Name,
		"latency":     latency,
		"mac_address": macAddress,
		"version":     initData.LocalVersion,
	}
	// 获取后端地址
	baseUrl := gcache.MustGet(context.Background(), "BackendBaseUrl").String()
	_, err = g.Client().Post(context.Background(), baseUrl+g_consts.PingBackendUrl, params)
	if err != nil {
		glog.Warning(context.Background(), "推送延迟到服务器时发生错误:", err)
		return
	}
	return
}
