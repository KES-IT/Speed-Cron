package push_utils

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-network-watcher/internal/global/g_consts"
)

type uCoffeePush struct{}

var CoffeePush = &uCoffeePush{}

// CoffeePushDataCore
//
//	@dc: 推送咖啡云信息-缓存
//	@params:
//	@response:
//	@author: laixin   @date:2023/6/6 14:15:14
func (u *uCoffeePush) CoffeePushDataCore() (err error) {
	coffeeCache, err := gcache.Get(context.TODO(), g_consts.CoffeeCacheKey)
	if err != nil {
		return err
	}
	if !coffeeCache.IsMap() {
		err = gerror.New("coffeeCache is not map")
		return err
	}
	err = u.PushCoffeeToBark(coffeeCache.Map())
	if err != nil {
		return err
	}
	return
}

// PushCoffeeToBark
//
//	@dc:
//	@params:
//	@response:
//	@author: laixin   @date:2023/6/6 14:26:09
func (u *uCoffeePush) PushCoffeeToBark(coffeeData g.Map) (err error) {
	var (
		url        = g_consts.PushCoreUrl
		pushClient = g.Client()
	)
	// 设置推送Header
	pushClient.SetHeader("Push-Sign", "Coffee_Push")
	// 设置为json
	pushClient.SetHeader("Content-Type", "application/json")
	response, err := pushClient.ContentJson().Post(context.Background(), url, coffeeData)
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			glog.Warning(context.Background(), "关闭响应失败", gconv.String(err))
		}
	}(response)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		err := gerror.New("状态码为:" + gconv.String(response.StatusCode) + response.ReadAllString())
		return err
	}
	return
}
