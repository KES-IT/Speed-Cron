package net_utils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net/http"
	"time"
)

type uNetUtils struct{}

var NetUtils = &uNetUtils{}

// HttpsLatency
//
//	@dc: 测试https延迟
//	@params:
//	@response:
//	@author: Administrator   @date:2023-06-17 14:01:06
func (u *uNetUtils) HttpsLatency() (latency int, err error) {
	url := "https://www.ithome.com/"
	start := time.Now()
	resp, err := http.Get(url)
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
//	@params:
//	@response:
//	@author: Administrator   @date:2023-06-17 14:03:41
func (u *uNetUtils) CoreLatency(initData g.Map) (err error) {
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
//	@params:
//	@response:
//	@author: Administrator   @date:2023-06-17 14:02:59
func (u *uNetUtils) PushLatencyToServer(initData g.Map, latency int) (err error) {
	url := "http://120.24.211.49:10441/UploadPingData"
	params := g.Map{
		"department": initData["department"],
		"staff_name": initData["name"],
		"latency":    latency,
	}
	_, err = g.Client().Post(context.Background(), url, params)
	if err != nil {
		glog.Warning(context.Background(), "推送延迟到服务器时发生错误:", err)
		return
	}
	return
}
