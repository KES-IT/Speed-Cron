package g_test

import (
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/test/gtest"
	_ "kes-cron/internal/packed"
	"testing"
)

// 测试解压文件
func Test_GDumpFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := gres.Export("speedCLI/speedtest.exe", "speed_cli/")
		t.Assert(err, nil)
	})
}
