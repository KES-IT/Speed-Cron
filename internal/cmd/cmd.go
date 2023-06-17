package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"kes-network-watcher/internal/boot"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gres.Dump()
			s := g.Server()
			// 初始化
			if err := boot.Boot(); err != nil {
				glog.Fatal(ctx, "初始化任务失败: ", err)
			}

			s.Run()
			return nil
		},
	}
)
