package main

import (
	"flag"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"kes-network-watcher/internal/cmd"
	_ "kes-network-watcher/internal/logic"
	_ "kes-network-watcher/internal/packed"
	binInfo "kes-network-watcher/utility/bin_utils"
	"os"
)

func main() {
	v := flag.Bool("v", false, "Show bin info.")
	flag.Parse()
	if *v {
		_, _ = fmt.Fprint(os.Stderr, binInfo.StringifyMultiLine())
		os.Exit(1)
	}
	cmd.Main.Run(gctx.New())
}
