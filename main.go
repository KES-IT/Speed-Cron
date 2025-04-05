package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"kes-cron/internal/cmd"
	"kes-cron/internal/global/g_consts"
	_ "kes-cron/internal/packed"
)

// GitTag 为编译时传入的版本号
// BackendBaseUrl 为编译时传入的后端地址
var (
	GitTag         = "unknown"
	BackendBaseUrl = "http://127.0.0.1:10441"
)

func main() {
	// 传入 GitTag 作为版本号
	cmd.LocalVersion = GitTag
	// 传入 BaseUrl 作为后端地址
	g_consts.BaseUrl = BackendBaseUrl

	_ = cmd.Main.AddCommand(cmd.Version)
	cmd.Main.Run(gctx.New())
}
