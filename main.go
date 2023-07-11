package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"kes-cron/internal/cmd"
	_ "kes-cron/internal/packed"
)

func main() {
	_ = cmd.Main.AddCommand(cmd.Version)
	cmd.Main.Run(gctx.New())
}
