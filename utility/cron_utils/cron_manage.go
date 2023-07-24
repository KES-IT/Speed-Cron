package cron_utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
	"kes-cron/internal/global/g_cache"
	"kes-cron/internal/global/g_consts"
	"kes-cron/internal/global/g_structs"
	"kes-cron/utility/cli_utils"
	"kes-cron/utility/net_utils"
	"reflect"
	"time"
)

type uCronManage struct{}

var CronManage = &uCronManage{}

// GetConfigAndStart
//
//	@dc: 获取定时任务管理器配置并启动
//	@author: hamster   @date:2023/6/20 13:41:26
func (u *uCronManage) GetConfigAndStart(ctx context.Context, initData *g_structs.InitData) (err error) {
	glog.Debug(ctx, "开始获取定时任务管理器配置")
	// 设备认证
	err = Auth.DeviceAuth(initData)
	if err != nil {
		glog.Warning(ctx, "设备认证失败: ", err)
		return err
	}
	glog.Debug(ctx, "设备认证成功")
	// 获取config
	speedInterval, pingInterval, cronStatus, err := getConfig()
	if err != nil {
		return
	}
	if cronStatus == 0 {
		glog.Warning(ctx, "定时任务管理器已关闭")
		removeAllCron()
		return
	}
	// 获取延迟检测定时任务
	localHTTPSCron := gcron.Search("HTTPS-Cron")
	if localHTTPSCron == nil {
		glog.Notice(ctx, "本地不存在定时任务,添加HTTPS-Cron定时任务")
		err = addPingCron(ctx, initData, pingInterval)
		if err != nil {
			glog.Warning(ctx, "添加HTTPS-Cron定时器失败")
			return
		}
	} else {
		HTTPSEntryPattern := reflect.ValueOf(localHTTPSCron).Elem().FieldByName("schedule").Elem().FieldByName("pattern")
		if HTTPSEntryPattern.IsValid() {
			if HTTPSEntryPattern.String() != pingInterval {
				glog.Notice(ctx, "更新HTTPS-Cron定时器")
				// 删除旧定时任务
				gcron.Stop("HTTPS-Cron")
				gcron.Remove("HTTPS-Cron")
				// 更新定时任务
				err = addPingCron(ctx, initData, pingInterval)
				if err != nil {
					glog.Warning(ctx, "更新HTTPS-Cron定时器失败")
					return
				}
				return
			} else {
				glog.Notice(ctx, "HTTPS-Cron定时器无需更新")
			}
		} else {
			glog.Warning(ctx, "HTTPS-Cron定时器无效")
		}

	}
	// 获取测速定时任务
	localSpeedCron := gcron.Search("Speed-Cron")
	if localSpeedCron == nil {
		glog.Notice(ctx, "本地不存在定时任务,添加Speed-Cron定时任务")
		err = addSpeedCron(ctx, initData, speedInterval)
		if err != nil {
			glog.Warning(ctx, "添加Speed-Cron定时器失败")
			return
		}
	} else {
		speedEntryPattern := reflect.ValueOf(localSpeedCron).Elem().FieldByName("schedule").Elem().FieldByName("pattern")
		if speedEntryPattern.String() != speedInterval {
			glog.Notice(ctx, "更新Speed-Cron定时器")
			// 删除旧定时任务
			gcron.Stop("Speed-Cron")
			gcron.Remove("Speed-Cron")
			// 更新定时任务
			err = addSpeedCron(ctx, initData, speedInterval)
			if err != nil {
				glog.Warning(ctx, "更新Speed-Cron定时器失败")
				return
			}
			return
		} else {
			glog.Notice(ctx, "Speed-Cron定时器无需更新")
		}
	}
	return
}

// removeAllCron 移除所有定时任务
func removeAllCron() {
	localSpeedCron := gcron.Search("Speed-Cron")
	if localSpeedCron != nil {
		glog.Debug(context.TODO(), "开始移除测速定时任务")
		gcron.Stop("Speed-Cron")
		gcron.Remove("Speed-Cron")
		glog.Debug(context.TODO(), "移除测速定时任务成功")
	}
	localHTTPSCron := gcron.Search("HTTPS-Cron")
	if localHTTPSCron != nil {
		glog.Debug(context.TODO(), "开始移除延迟检测定时任务")
		gcron.Stop("HTTPS-Cron")
		gcron.Remove("HTTPS-Cron")
		glog.Debug(context.TODO(), "移除延迟检测定时任务成功")
	}

	return
}

// getConfig 获取定时任务管理器配置
func getConfig() (speedInterval string, pingInterval string, cronStatus int, err error) {
	// 获取Mac地址
	_, macAddress := net_utils.NetworkInfo.GetMacAddress()
	// 获取配置
	glog.Debug(context.TODO(), "当前mac地址为", macAddress)
	// 获取后端地址
	baseUrl := gcache.MustGet(context.Background(), "BackendBaseUrl").String()
	response, err := g.Client().SetTimeout(5*time.Second).Post(context.TODO(), baseUrl+g_consts.ConfigBackendUrl, g.Map{
		"mac_address": macAddress,
	})
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			glog.Warning(context.TODO(), "getConfig 关闭response失败: ", err)
		}
	}(response)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		err = gerror.Newf("获取配置失败，错误码：%d", response.StatusCode)
		return
	}
	// 解析配置
	configMap := gjson.New(response.ReadAllString())
	// 获取测速间隔与延迟检测间隔
	speedInterval = configMap.Get("data.speed_interval").String()
	pingInterval = configMap.Get("data.ping_interval").String()
	cronStatus = configMap.Get("data.cron_status").Int()
	return
}

// addSpeedCron 添加测速定时任务
func addSpeedCron(ctx context.Context, initData *g_structs.InitData, timePattern string) (err error) {
	glog.Notice(ctx, "开始定时测速服务", timePattern)
	_, err = gcron.AddSingleton(ctx, timePattern, func(ctx context.Context) {
		// 判断是否在更新中
		if gcache.MustGet(ctx, g_cache.UpdateCacheKey).Bool() {
			glog.Warning(ctx, "正在更新客户端程序，跳过本次测速")
			return
		}
		_ = gcache.Set(ctx, g_cache.SpeedCacheKey, true, 0)

		err := cli_utils.CmdCore.StartSpeedCmd(ctx, initData)
		if err != nil {
			glog.Error(ctx, "定时测速服务失败: ", err)
			return
		}

		_, _ = gcache.Remove(ctx, g_cache.SpeedCacheKey)
	}, "Speed-Cron")
	if err != nil {
		glog.Warning(ctx, "添加定时测速服务失败: ", err)
		return err
	}
	return
}

// addPingCron 添加延迟检测定时任务
func addPingCron(ctx context.Context, initData *g_structs.InitData, timePattern string) (err error) {
	glog.Notice(ctx, "开始HTTPS延迟定时检测服务", timePattern)
	_, err = gcron.AddSingleton(ctx, timePattern, func(ctx context.Context) {
		err := net_utils.NetUtils.CoreLatency(ctx, initData)
		if err != nil {
			glog.Error(ctx, "HTTPS延迟检测失败: ", err)
			return
		}
	}, "HTTPS-Cron")
	if err != nil {
		glog.Warning(ctx, "添加HTTPS延迟定时检测服务失败: ", err)
		return
	}
	return
}
