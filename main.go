package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"kes-cron/internal/cmd"
	_ "kes-cron/internal/packed"
)

// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown

var (
	GitTag = "cool"
)

func main() {
	// 传入 GitTag 作为版本号
	cmd.GitTag = GitTag
	_ = cmd.Main.AddCommand(cmd.Version)
	cmd.Main.Run(gctx.New())
}
