package cli_utils

import (
	"bufio"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-cron/internal/global/g_consts"
	"time"
)

type uCmdCore struct{}

var CmdCore = &uCmdCore{}

// StartSpeedCmd
//
//	@dc: 启动speedtest命令
//	@author: hamster   @date:2023/6/20 10:06:06
func (u *uCmdCore) StartSpeedCmd(ctx context.Context, initData *g_consts.InitData) (err error) {
	_ = gcache.Set(ctx, "speedtest", true, 1*time.Minute)
	cmd := CliUtils.CreateSpeedCmd()
	if cmd == nil {
		glog.Warning(ctx, "创建命令失败,获取测速节点失败")
		err = gerror.New("创建命令失败,获取测速节点失败")
		return
	}
	// 获取命令的标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		glog.Warning(ctx, "获取标准输出管道时发生错误:", err)
		return
	}
	// 启动命令
	err = cmd.Start()
	if err != nil {
		glog.Warning(ctx, "启动命令时发生错误:", err)
		return
	}
	var (
		scanner             = bufio.NewScanner(stdout)
		defaultBars         = Bar.InitDefaultBar()
		uploadNetDataStruct = &NetInfoUploadData{}
	)
	uploadNetDataStruct.Department = gconv.String(initData.Department)
	uploadNetDataStruct.StaffName = gconv.String(initData.Name)
	// 持续获取输出
	for scanner.Scan() {
		// 获取输出行
		line := scanner.Bytes()
		ok, err := CmdProgress.CmdCoreProgress(ctx, string(line), defaultBars, uploadNetDataStruct)
		if err != nil {
			glog.Warning(ctx, "处理命令行输出时发生错误:", err)
			_ = cmd.Process.Kill()
			return err
		}
		if ok {
			glog.Debug(ctx, "单次speedtest测速处理已完成")
			break
		}
	}
	// 进行speedtestCLI的退出
	glog.Debug(ctx, "关闭speedtestCLI中", cmd.Process.Pid)
	err = cmd.Process.Kill()
	if err != nil {
		glog.Warning(ctx, "退出speedtestCLI时发生错误:", err)
	} else {
		glog.Debug(ctx, "speedtestCLI已退出", cmd.Process.Pid)
	}
	_, _ = gcache.Remove(ctx, "speedtest")
	return
}
