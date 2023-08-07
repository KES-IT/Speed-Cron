package cli_utils

import (
	"bufio"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"kes-cron/internal/global/g_cache"
	"kes-cron/internal/global/g_structs"
)

type uCmdCore struct{}

var CmdCore = &uCmdCore{}

// StartSpeedCmd
//
//	@dc: 启动SpeedTest命令
//	@author: hamster   @date:2023/6/20 10:06:06
func (u *uCmdCore) StartSpeedCmd(ctx context.Context, initData *g_structs.InitData) (err error) {
	// 创建命令
	cmd := CliUtils.CreateSpeedCmd()
	if cmd == nil {
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
		uploadNetDataStruct = &NetInfoUploadData{
			Department: initData.Department,
			StaffName:  initData.Name,
		}
	)

	// 持续获取输出
	for scanner.Scan() {
		// 获取更新核心状态，判断是否即将更新
		if gcache.MustGet(ctx, g_cache.UpdateCacheKey).Bool() {
			glog.Warning(ctx, "正在更新客户端程序，中止测速服务")
			break
		}
		ok, err := CmdProgress.CmdCoreProgress(ctx, string(scanner.Bytes()), defaultBars, uploadNetDataStruct)
		if err != nil {
			glog.Warning(ctx, "处理命令行输出时发生错误:", err)
			_ = cmd.Process.Kill()
			return err
		}
		if ok {
			glog.Debug(ctx, "单次SpeedTest测速处理已完成")
			break
		}
	}
	// 进行SpeedTestCLI的退出
	glog.Debug(ctx, "关闭SpeedTestCLI中", cmd.Process.Pid)
	err = cmd.Process.Kill()
	if err != nil {
		glog.Warning(ctx, "退出SpeedTestCLI时发生错误:", err)
	} else {
		glog.Debug(ctx, "SpeedTestCLI已退出", cmd.Process.Pid)
	}
	return
}
