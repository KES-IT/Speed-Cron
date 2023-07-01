package cli_utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
	"kes-cron/internal/global/g_consts"
	"kes-cron/utility/net_utils"
	"os/exec"
	"time"
)

type uCliUtils struct{}

var CliUtils = &uCliUtils{}

func (s *uCliUtils) CreateSpeedCmd() *exec.Cmd {
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
		return nil
	}
	if response.StatusCode != 200 {
		err = gerror.Newf("获取配置失败，错误码：%d", response.StatusCode)
		return nil
	}
	// 解析配置
	configMap := gjson.New(response.ReadAllString())

	// 配置测速节点
	serverIdCmd := "--server-id=" + configMap.Get("data.speed_server_id").String()
	glog.Debug(context.Background(), "测速节点:", serverIdCmd)
	return exec.Command("speed_cli/speedCLI/speedtest.exe", "--accept-gdpr", "--accept-license", serverIdCmd,
		"--progress=yes", "--format=json", "--progress-update-interval=200")
}

/*// StartSingleSpeedTest
//
//	@dc: 启动单次测速
//	@params:
//	@response:
//	@author: Administrator   @date:2023-06-17 10:37:40
func (u *uCliUtils) StartSingleSpeedTest(initData g.Map) (err error) {
	err = gres.Export("speed/speed_bin.exe", "speed_cli/")
	if err != nil {
		glog.Warning(context.Background(), "导出资源文件时发生错误:", err)
		return
	}
	departmentCmd := "-department=" + gconv.String(initData["department"])
	nameCmd := "-name=" + gconv.String(initData["name"])
	glog.Info(context.Background(), "部门:", departmentCmd, "姓名:", nameCmd)
	cmd := exec.Command("speed_cli/speed/speed_bin.exe", departmentCmd, nameCmd)
	// 启动命令
	glog.Notice(context.Background(), "启动测速命令")
	err = cmd.Start()
	if err != nil {
		glog.Warning(context.Background(), "启动命令时发生错误:", err)
		return
	}
	glog.Info(context.Background(), "运行pid为:", cmd.Process.Pid)
	status, err := cmd.Process.Wait()
	if status.Exited() {
		glog.Info(context.Background(), "单次测速执行完成")
	}
	return
}*/
