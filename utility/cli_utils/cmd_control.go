package cli_utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
	"kes-cron/internal/global/g_consts"
	"kes-cron/utility/net_utils"
	"os/exec"
	"runtime"
	"time"
)

type uCliUtils struct{}

var CliUtils = &uCliUtils{}

// CreateSpeedCmd 创建测速命令
func (s *uCliUtils) CreateSpeedCmd() *exec.Cmd {
	// 获取Mac地址
	_, macAddress := net_utils.NetworkInfo.GetMacAddress()
	// 获取配置
	glog.Debug(context.TODO(), "重新根据mac获取配置信息", macAddress)
	// 从服务器中获取配置
	response, err := g.Client().SetTimeout(5*time.Second).Post(context.TODO(), g_consts.BackendBaseUrl()+g_consts.ConfigBackendUrl, g.Map{
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
		glog.Warningf(context.Background(), "获取配置失败，错误码：%d", response.StatusCode)
		return nil
	}
	// 解析配置
	configMap := gjson.New(response.ReadAllString())

	// 配置测速节点
	serverIdCmd := "--server-id=" + configMap.Get("data.speed_server_id").String()

	if configMap.Get("data.speed_server_id").String() == "" {
		serverIdCmd = ""
	}
	glog.Debug(context.Background(), "测速节点:", serverIdCmd)
	speedtestCliPath := ""

	if runtime.GOOS == "windows" {
		speedtestCliPath = "speed_cli/speedCLI/speedtest-win.exe"
	}
	if runtime.GOOS == "linux" {
		if runtime.GOARCH == "amd64" {
			speedtestCliPath = "speed_cli/speedCLI/speedtest-linux-amd64"
		} else {
			speedtestCliPath = "speed_cli/speedCLI/speedtest-linux-arm64"
		}
	}
	if runtime.GOOS == "darwin" {
		speedtestCliPath = "speed_cli/speedCLI/speedtest-mac"
	}

	return exec.Command(speedtestCliPath, "--accept-gdpr", "--accept-license", serverIdCmd,
		"--progress=yes", "--format=json", "--progress-update-interval=500")
}
