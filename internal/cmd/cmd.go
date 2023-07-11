package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"kes-cron/internal/boot"
	"kes-cron/utility/cron_utils"
)

var (
	Version = &gcmd.Command{
		Name:        "version",
		Brief:       "return version",
		Description: "return version",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			glog.Print(ctx, "v0.0.4")
			return nil
		},
	}
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gres.Dump()
			err = gres.Export("speedCLI/speedtest.exe", "speed_cli/")
			if err != nil {
				glog.Warning(ctx, "导出资源文件时发生错误:", err)
				return err
			}
			initData := g.Map{
				"department": parser.GetOpt("department").String(),
				"name":       parser.GetOpt("name").String(),
			}
			glog.Notice(ctx, "当前部门: ", initData["department"], " 当前员工: ", initData["name"])
			if initData["department"] == "" || initData["name"] == "" {
				glog.Warning(ctx, "初始化任务失败: ", "参数错误", "设置为默认值")
				initData["department"] = "未知部门"
				initData["name"] = "未知员工"
			}
			err = cron_utils.Auth.DeviceAuth(initData)
			if err != nil {
				glog.Warning(ctx, "设备认证失败: ", err)
				return err
			}

			// 获取设备信息
			serverInitData, err := cron_utils.Auth.GetDeviceInfo()
			if err != nil {
				glog.Warning(ctx, "获取设备信息失败: ", err)
				return err
			}

			// 初始化
			if err := boot.Boot(serverInitData); err != nil {
				glog.Fatal(ctx, "初始化任务失败: ", err)
			}
			// 启动服务
			s := g.Server()
			s.Run()
			return nil
		},
	}
)
