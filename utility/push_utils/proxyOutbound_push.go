package push_utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-network-watcher/internal/global/g_consts"
)

// GetTotalOutbound
//
//	@dc: 获取总出口流量
//	@params:
//	@response:
//	@author: laixin   @date:2023/5/11 00:20:58
func (u *uPushUtils) GetTotalOutbound() (outBoundStr string, err error) {
	var (
		totalOutbound float64
	)
	proxyUserList, err := PushUtils.GetProxyUser()
	if err != nil {
		return
	}
	// 计算全部用户的出口流量
	for _, user := range proxyUserList {
		userMap := gconv.Map(user)
		totalOutbound += gconv.Float64(userMap["up"]) + gconv.Float64(userMap["down"])
	}
	// byte -> G
	outBoundG := totalOutbound / g_consts.ByteToG
	outBoundStr = fmt.Sprintf("%.3f", outBoundG)
	return
}

// StoreOutbound
//
//	@dc: 存储出口流量
//	@params:
//	@response:
//	@author: laixin   @date:2023/5/11 02:20:10
func (u *uPushUtils) StoreOutbound() (err error) {
	var (
		nowTime = gtime.Now().String()
	)
	outBoundStr, err := PushUtils.GetTotalOutbound()
	if err != nil {
		return
	}
	outBoundInfo := g.Map{
		"time":     nowTime,
		"outBound": outBoundStr,
	}
	err = gcache.Set(context.Background(), g_consts.ProxyOutboundCacheKey, outBoundInfo, 0)
	if err != nil {
		return err
	}
	return
}

// GetUsedOutboundAndPush
//
//	@dc: 获取已使用的出口流量
//	@params:
//	@response:
//	@author: laixin   @date:2023/5/11 02:22:34
func (u *uPushUtils) GetUsedOutboundAndPush() (err error) {
	if outBoundData, err := gcache.Get(context.Background(), g_consts.ProxyOutboundCacheKey); err != nil {
		if err != nil {
			return err
		} else if !outBoundData.IsMap() {
			err = fmt.Errorf("outBoundData is not map")
			return err
		}
	} else if !outBoundData.IsEmpty() {
		outBoundStr, err := u.GetTotalOutbound()
		outBoundMap := outBoundData.MapStrStr()
		duringTime := "从" + gtime.NewFromStr(outBoundMap["time"]).Format("G:i") + "到" + gtime.Now().Format("G:i") + " " + gtime.NewFromStr(gtime.Now().String()).Sub(gtime.NewFromStr(outBoundMap["time"])).String()
		usedOutBound := fmt.Sprintf("%.2f", gconv.Float64(outBoundStr)-gconv.Float64(outBoundMap["outBound"]))
		// 推送到Bark
		err = PushUtils.PushOutboundToBark(usedOutBound, duringTime)
		if err != nil {
			return err
		}
	}
	return
}

// PushOutboundToBark
//
//	@dc: 推送出口流量到Bark
//	@params:
//	@response:
//	@author: laixin   @date:2023/5/11 02:37:19
func (u *uPushUtils) PushOutboundToBark(usedOutBound, duringTime string) (err error) {
	var (
		url        = g_consts.PushCoreUrl
		pushClient = g.Client()
	)
	pushClient.SetHeader("Push-Sign", "ProxyOutbound_Push")
	pushClient.SetHeader("Content-Type", "application/json")

	response, err := pushClient.ContentJson().Post(context.Background(), url, g.Map{
		"duringTime":   duringTime,
		"usedOutBound": usedOutBound,
	})
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(response)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		err := gerror.New("推送失败")
		return err
	}
	return
}
