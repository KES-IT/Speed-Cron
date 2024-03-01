package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"kes-cron/internal/boot"
	"kes-cron/internal/global/g_consts"
	"kes-cron/internal/global/g_structs"
	"kes-cron/utility/cron_utils"
	"time"
)

var (
	LocalVersion = "unknown"
)

func GetVersion() string {
	return LocalVersion
}

var (
	Version = &gcmd.Command{
		Name:        "version",
		Brief:       "return version",
		Description: "show exe version",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			fmt.Print(GetVersion())
			return nil
		},
	}

	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start kes-speed-cron client",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gres.Dump()
			err = gres.Export("speedCLI/speedtest.exe", "speed_cli/")
			if err != nil {
				glog.Warning(ctx, "导出资源文件时发生错误:", err)
				return err
			}
			// 判断后端地址是否为空
			if g_consts.BackendBaseUrl() == "" {
				glog.Warning(ctx, "后端地址为空: ", "请检查下载渠道")
				time.Sleep(5 * time.Second)
				return nil
			}
			// 初始化数据
			initData := &g_structs.InitData{
				Department: parser.GetOpt("department").String(),
				Name:       parser.GetOpt("name").String(),
			}
			glog.Notice(ctx, "当前初始化部门: ", initData.Department, " 当前初始化员工: ", initData.Name)
			if initData.Department == "" || initData.Name == "" {
				glog.Warning(ctx, "初始化任务失败: ", "参数错误", "设置为默认值")
				initData.Department = "未知部门-吉他维修部"
				initData.Name = "未知员工-方大同"
			}

			// 第一次设备注册\认证
			err = cron_utils.Auth.DeviceAuth(initData)
			if err != nil {
				glog.Warning(ctx, "设备认证失败: ", err)
				time.Sleep(5 * time.Second)
				return err
			}

			// 获取设备信息
			serverInitData, err := cron_utils.Auth.GetDeviceInfo()
			if err != nil {
				glog.Warning(ctx, "获取设备信息失败: ", err)
				time.Sleep(5 * time.Second)
				return err
			}

			// 设置本地版本号
			serverInitData.LocalVersion = GetVersion()
			glog.Notice(ctx, "当前客户端版本: ", GetVersion())

			// 设置后端地址
			glog.Notice(ctx, "当前后端地址: ", g_consts.BackendBaseUrl())

			// 初始化
			if err := boot.Boot(serverInitData); err != nil {
				glog.Fatal(ctx, "初始化任务失败: ", err)
				time.Sleep(5 * time.Second)
				return err
			}

			// 启动服务
			s := g.Server()
			s.Run()
			return nil
		},
	}
)
