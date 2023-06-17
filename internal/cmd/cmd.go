package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"kes-cron/internal/boot"
	"time"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gres.Dump()
			initData := g.Map{
				"department": parser.GetOpt("department").String(),
				"name":       parser.GetOpt("name").String(),
			}
			glog.Notice(ctx, "当前部门: ", initData["department"], " 当前员工: ", initData["name"])
			if initData["department"] == "" || initData["name"] == "" {
				glog.Warning(ctx, "初始化任务失败: ", "参数错误")
				time.Sleep(5 * time.Second)
				return err
			}
			s := g.Server()
			// 初始化
			if err := boot.Boot(initData); err != nil {
				glog.Fatal(ctx, "初始化任务失败: ", err)
			}

			s.Run()
			return nil
		},
	}
)
