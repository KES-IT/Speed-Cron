package cli_utils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/util/gconv"
	"os/exec"
)

type uCliUtils struct{}

var CliUtils = &uCliUtils{}

// StartSingleSpeedTest
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
}
