package network_utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-network-watcher/internal/global/g_consts"
)

type uCoffee struct{}

var Coffee = &uCoffee{}

// GetCoffeeInfo
//
//	@dc: 获取咖啡云信息
//	@params:
//	@response:
//	@author: laixin   @date:2023/6/6 04:15:01
func (u *uCoffee) GetCoffeeInfo() (err error) {
	err, authData := getAuthData()
	if err != nil || authData == "" {
		return
	}
	coffeeClient := gclient.New()
	coffeeClient.SetHeaderMap(map[string]string{
		"Authorization": authData,
	})
	response, err := coffeeClient.Get(context.TODO(), g_consts.CoffeeBaseUrl, nil)
	if err != nil {
		return
	}
	infoData := gconv.Map(gconv.Map(response.ReadAllString())["data"])
	if infoData["d"] == nil || infoData["transfer_enable"] == nil {
		return
	}
	usedBound := gconv.Float64(infoData["d"]) / 1024 / 1024 / 1010
	planBound := gconv.Float64(infoData["transfer_enable"]) / 1024 / 1024 / 1024
	remainBound := planBound - usedBound
	// 保留两位小数
	usedBoundStr := fmt.Sprintf("%.2f", usedBound)
	planBoundStr := fmt.Sprintf("%.2f", planBound)
	remainBoundStr := fmt.Sprintf("%.2f", remainBound)
	coffeeCache := g.Map{
		"usedBound":   usedBoundStr,
		"planBound":   planBoundStr,
		"remainBound": remainBoundStr,
	}
	err = gcache.Set(context.TODO(), g_consts.CoffeeCacheKey, coffeeCache, 0)
	if err != nil {
		return err
	}
	return
}

func getAuthData() (err error, authData string) {
	url := g_consts.CoffeeLoginUrl
	response, err := gclient.New().Post(context.TODO(), url, g_consts.CoffeeAuthData)
	if err != nil {
		return
	}
	return nil, gjson.New(response.ReadAllString()).Get("data.auth_data").String()
}
