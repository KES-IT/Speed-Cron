package push_utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-network-watcher/internal/global/g_consts"
	"kes-network-watcher/utility/network_utils"
	"time"
)

type uPushUtils struct{}

var PushUtils = &uPushUtils{}

// GetProxyNetwork
//
//	@dc: 获取科学上网速度
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/19 18:47:05
func (u *uPushUtils) GetProxyNetwork() (proxyNetworkUp string, err error) {
	proxyData, err := gcache.Get(context.Background(), g_consts.ProxyNetworkCacheKey)
	if err != nil {
		return "", err
	}
	proxyNetworkUp = gjson.New(proxyData.Map()).Get("txSpeedMbps").String()
	return proxyNetworkUp, nil
}

// ProxyPushToBark
//
//	@dc: 代理网速推送到bark
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/19 18:47:47
func (u *uPushUtils) ProxyPushToBark(proxyNetworkUp, maxFlowUser string, maxFlow int) (err error) {
	var (
		url        = g_consts.PushCoreUrl
		pushClient = g.Client()
	)
	// 设置推送Header
	pushClient.SetHeader("Push-Sign", "ProxyNetwork_Push")
	// 设置为json
	pushClient.SetHeader("Content-Type", "application/json")
	response, err := pushClient.ContentJson().Post(context.Background(), url, g.Map{
		"proxyNetworkUp": proxyNetworkUp,
		"maxFlow":        maxFlow,
		"maxFlowUser":    maxFlowUser,
	})
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

// GetProxyUser
//
//	@dc: 获取代理占用用户
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/23 17:57:52
func (u *uPushUtils) GetProxyUser() (userList []interface{}, err error) {
	var (
		session map[string]string
		url     = g_consts.XuiBaseUrl + "/xui/inbound/list"
	)

	// 尝试获取缓存中的session
	sessionData, err := gcache.Get(context.Background(), g_consts.ProxySessionCacheKey)
	if err != nil {
		return nil, err
	}

	if sessionData.IsNil() {
		// 重新获取session
		session, err = network_utils.ProxyNetwork.GetXuiSession()
		if err != nil {
			return nil, err
		}
	} else {
		session = sessionData.MapStrStr()
	}

	// 获取代理占用用户
	post, err := g.Client().SetCookieMap(session).Post(context.Background(), url)
	defer func(post *gclient.Response) {
		err := post.Close()
		if err != nil {
			glog.Warning(context.Background(), "关闭响应失败", gconv.String(err))
		}
	}(post)
	if err != nil {
		return nil, err
	}
	if post.StatusCode != 200 {
		err = gerror.New("状态码为:" + gconv.String(post.StatusCode) + post.ReadAllString())
		return nil, err
	}
	jsonData := gjson.New(post.ReadAllString())
	userList = jsonData.Get("obj").Array()
	if len(userList) == 0 {
		err := gerror.New("代理用户为空")
		return nil, err
	}
	return userList, nil
}

// ProxyPushCore
//
//	@dc: 代理网速推送核心
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/19 18:48:13
func (u *uPushUtils) ProxyPushCore(ctx context.Context) (err error) {
	var (
		// 速率限制
		speedLimit = "6"
		limitTime  = 10
	)

	// 获取科学上网速度
	proxyNetworkUp, err := u.GetProxyNetwork()
	if err != nil {
		return
	}
	proxyNetworkUpSpeed := gconv.Float64(proxyNetworkUp)
	// 进行10s内超过速率限制次数判断
	if proxyNetworkUpSpeed > gconv.Float64(speedLimit) {
		// 速率超过限制
		// 获取缓存中的速率超过限制次数
		count, err := gcache.Get(ctx, g_consts.ProxyCountCacheKey)
		if err != nil {
			return err
		}
		if count == nil {
			err := gcache.Set(ctx, g_consts.ProxyCountCacheKey, 1, 20*time.Second)
			if err != nil {
				return err
			}
			// 获取占用用户
			userList, err := u.GetProxyUser()
			if err != nil {
				return err
			}
			// 获取用户当前流量存入缓存
			err = gcache.Set(ctx, g_consts.ProxyUserFlowCacheKey, userList, 0)
			if err != nil {
				return err
			}
			return nil
		}
		countInt := count.Int()
		if countInt > limitTime {
			// 清空缓存
			_, _ = gcache.Remove(ctx, g_consts.ProxyCountCacheKey)
			_, _ = gcache.Remove(ctx, g_consts.ProxyUserFlowCacheKey)
		}
		if countInt == limitTime {
			var (
				maxFLow     int
				maxFLowUser string
			)
			// 计算用户流量变化
			// 获取缓存中的用户流量
			userList, _ := gcache.Get(ctx, g_consts.ProxyUserFlowCacheKey)
			if userList == nil {
				glog.Warning(context.Background(), "未成功获取到用户流量")
			}
			if userList != nil {
				// 获取当前用户流量
				userListNow, err := u.GetProxyUser()
				if err != nil {
					glog.Warning(context.Background(), "未获取到当前用户流量")
					glog.Warning(context.Background(), err)
				} else if len(userListNow) != 0 {
					// 计算用户流量变化
					for _, user := range userList.Array() {
						userMap := gconv.Map(user)
						for _, userNow := range userListNow {
							userNowMap := gconv.Map(userNow)
							if userMap["id"] == userNowMap["id"] {
								// 计算用户流量变化
								totalFlow := gconv.Int(userNowMap["down"])/1024/1024 - gconv.Int(userMap["down"])/1024/1024
								if totalFlow >= maxFLow {
									maxFLow = totalFlow
									maxFLowUser = gconv.String(userNowMap["remark"])
								}
							}
						}
					}
				}

			}
			// 进行推送
			err = u.ProxyPushToBark(proxyNetworkUp, maxFLowUser, maxFLow)
			if err != nil {
				glog.Warning(ctx, "推送失败:"+err.Error())
			}
			glog.Notice(ctx, "推送成功")
			// 清空缓存
			_, err = gcache.Remove(ctx, g_consts.ProxyCountCacheKey)
			if err != nil {
				glog.Warning(ctx, "清空缓存失败:", err)
			}
			_, err = gcache.Remove(ctx, g_consts.ProxyUserFlowCacheKey)
			if err != nil {
				glog.Warning(ctx, "清空缓存失败:", err)
			}
		} else {
			// 速率超过限制次数+1
			err = gcache.Set(ctx, g_consts.ProxyCountCacheKey, countInt+1, gcache.MustGetExpire(ctx, g_consts.ProxyCountCacheKey))
			if err != nil {
				return err
			}
		}

	} else {
		// 速率未超过限制
	}

	return
}
