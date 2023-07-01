package cron_utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
	"kes-cron/internal/global/g_consts"
	"kes-cron/utility/net_utils"
	"time"
)

type uAuth struct{}

var Auth = &uAuth{}

// DeviceAuth
//
//	@dc: 设备认证
//	@params:
//	@response:
//	@author: hamster   @date:2023/6/20 11:42:38
func (u *uAuth) DeviceAuth(initData g.Map) (err error) {
	// 获取内网与mac地址
	internalIp, macAddress := net_utils.NetworkInfo.GetMacAddress()
	// 进行设备认证
	response, err := g.Client().SetTimeout(5*time.Second).Post(context.TODO(), g_consts.AuthBackendUrl, g.Map{
		"internal_ip": internalIp,
		"mac_address": macAddress,
		"department":  initData["department"],
		"staff_name":  initData["name"],
	})
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			glog.Error(context.Background(), "关闭response时发生错误:", err)
		}
	}(response)
	if err != nil {
		return err
	}
	// 获取认证结果
	if response.StatusCode != 200 {
		err = gerror.Newf("认证失败，错误码：%d", response.StatusCode)
		return
	}
	return
}

// GetDeviceInfo
//
//	@dc:
//	@params:
//	@response:
//	@author: hamster   @date:2023/6/20 17:34:24
func (u *uAuth) GetDeviceInfo() (initData g.Map, err error) {
	// 获取Mac地址
	_, macAddress := net_utils.NetworkInfo.GetMacAddress()
	// 获取配置
	glog.Debug(context.TODO(), "重新根据mac获取配置信息", macAddress)
	response, err := g.Client().SetTimeout(5*time.Second).Post(context.TODO(), g_consts.ConfigBackendUrl, g.Map{
		"mac_address": macAddress,
	})
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			glog.Error(context.Background(), "关闭response时发生错误:", err)
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
	initData = g.Map{
		"department": configMap.Get("data.department").String(),
		"name":       configMap.Get("data.staff_name").String(),
	}
	glog.Debug(context.TODO(), "获取到的当前个人信息为", "部门", initData["department"], "姓名", initData["name"])
	return
}
