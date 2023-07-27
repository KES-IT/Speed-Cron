package net_utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-cron/internal/global/g_consts"
	"kes-cron/internal/global/g_structs"
	"time"
)

type uNetUtils struct{}

var NetUtils = &uNetUtils{}

// HttpsLatency
//
//	@dc: 测试https延迟
//	@author: Hamster   @date:2023-06-17 14:01:06
func (u *uNetUtils) HttpsLatency(url string) (latency int, err error) {
	start := time.Now()
	resp, err := g.Client().Timeout(5*time.Second).Get(context.Background(), url)
	if err != nil {
		glog.Warning(context.Background(), "请求出错:", err)
		return
	}
	defer func(resp *gclient.Response) {
		err = resp.Close()
		if err != nil {
			glog.Warning(context.Background(), "HttpsLatency 关闭请求时发生错误:", err)
		}
	}(resp)
	latency = gconv.Int(time.Since(start).Milliseconds())
	return latency, nil
}

// CoreLatency
//
//	@dc: 测试延迟核心服务
//	@author: Hamster   @date:2023-06-17 14:03:41
func (u *uNetUtils) CoreLatency(ctx context.Context, initData *g_structs.InitData) (err error) {
	latency, err := u.HttpsLatency(g_consts.PingUrl)
	if err != nil {
		glog.Warning(ctx, "请求出错:", err)
		return
	}
	glog.Info(ctx, "HTTPS延迟:", latency)
	err = u.PushLatencyToServer(initData, latency)
	if err != nil {
		glog.Warning(ctx, "推送延迟到服务器时发生错误:", err)
		return err
	}
	glog.Info(ctx, "开始多节点延迟测试")
	err = u.MultiWebsiteLatencyCore()
	if err != nil {
		glog.Warning(ctx, "多节点延迟核心服务时发生错误:", err)
		return err
	}
	glog.Info(ctx, "多节点延迟测试完成")
	return
}

// MultiWebsiteLatencyCore
//
//	@dc: 多节点延迟核心服务
//	@author: laixin   @date:2023/7/20 17:01:51
func (u *uNetUtils) MultiWebsiteLatencyCore() (err error) {
	monitorList, err := u.GetMonitorList()
	if err != nil {
		glog.Warning(context.Background(), "获取监控列表时发生错误:", err)
		return
	}

	_, macAddress := NetworkInfo.GetMacAddress()

	// 遍历监控列表进行延迟测试
	for _, monitor := range monitorList {
		monitorJson := gjson.New(monitor)
		websiteUrl := monitorJson.Get("website_url").String()
		glog.Info(context.Background(), "开始测试: ", websiteUrl)
		latency, httpErr := u.HttpsLatency(websiteUrl)
		httpErrStr := ""
		if httpErr != nil {
			glog.Warning(context.Background(), "请求出错:", err)
			httpErrStr = httpErr.Error()
		} else {
			glog.Info(context.Background(), websiteUrl+" HTTPS延迟: ", latency)
		}
		// 推送延迟到服务器
		// 获取后端地址
		baseUrl := gcache.MustGet(context.Background(), "BackendBaseUrl").String()
		monitorRes, err := g.Client().Post(context.Background(), baseUrl+g_consts.MonitorLogBackendUrl, g.Map{
			"mac_address": macAddress,
			"website_id":  monitorJson.Get("id").Int(),
			"website_url": websiteUrl,
			"latency":     latency,
			"err_msg":     httpErrStr,
		})
		if err != nil {
			glog.Warning(context.Background(), "推送延迟到服务器时发生错误:", err)
		}
		// 关闭请求
		if monitorRes != nil {
			err = monitorRes.Close()
			if err != nil {
				glog.Warning(context.Background(), "MultiWebsiteLatencyCore 关闭请求时发生错误:", err)
			}
		}
	}
	return
}

// GetMonitorList
//
//	@dc: 获取监控列表
//	@author: laixin   @date:2023/7/20 16:55:37
func (u *uNetUtils) GetMonitorList() (monitorList []interface{}, err error) {
	// 获取后端地址
	baseUrl := gcache.MustGet(context.Background(), "BackendBaseUrl").String()
	monitorListRes, err := g.Client().Timeout(5*time.Second).Get(context.Background(), baseUrl+g_consts.MonitorListBackendUrl)
	defer func(monitorListRes *gclient.Response) {
		err := monitorListRes.Close()
		if err != nil {
			glog.Warning(context.Background(), "GetMonitorList 关闭请求时发生错误:", err)
		}
	}(monitorListRes)
	if err != nil {
		glog.Warning(context.Background(), "获取监控列表时发生错误:", err)
		return
	}
	monitorList = gjson.New(monitorListRes.ReadAllString()).Get("data.website_list").Array()
	if len(monitorList) == 0 {
		glog.Warning(context.Background(), "获取监控列表时发生错误: 监控列表为空")
		err = gerror.New("获取监控列表时发生错误: 监控列表为空")
		return
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
	pushRes, err := g.Client().Post(context.Background(), baseUrl+g_consts.PingBackendUrl, params)
	defer func(pushRes *gclient.Response) {
		err := pushRes.Close()
		if err != nil {
			glog.Warning(context.Background(), "PushLatencyToServer 关闭请求时发生错误:", err)
		}
	}(pushRes)
	if err != nil {
		glog.Warning(context.Background(), "推送延迟到服务器时发生错误:", err)
		return
	}
	return
}
