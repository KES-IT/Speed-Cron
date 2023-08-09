package g_test

import (
	"context"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/test/gtest"
	_ "kes-cron/internal/packed"
	"kes-cron/utility/cli_utils"
	"kes-cron/utility/net_utils"
	"kes-cron/utility/update_utils"
	"testing"
)

// 测试解压文件
func Test_GDumpFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := gres.Export("speedCLI/speedtest.exe", "speed_cli/")
		t.Assert(err, nil)
	})
}

// 测试模拟单次测速
func Test_Speed_Single(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := cli_utils.CmdCore.StartSpeedCmd(context.Background(), InitData)
		t.Assert(err, nil)
	})
}

// 测试多站点延迟测试
func Test_Website_Latency(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := net_utils.NetUtils.CoreLatency(context.Background(), InitData)
		t.Assert(err, nil)
	})
}

// 测试自动更新模块
func Test_Auto_Update(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := update_utils.AutoUpdate.UpdateCore(context.Background(), InitData)
		t.Assert(err, nil)
	})
}
