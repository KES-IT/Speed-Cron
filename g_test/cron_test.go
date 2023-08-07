package g_test

import (
	"context"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/test/gtest"
	"kes-cron/internal/global/g_consts"
	"kes-cron/internal/global/g_structs"
	_ "kes-cron/internal/packed"
	"kes-cron/utility/cli_utils"
	"kes-cron/utility/net_utils"
	"kes-cron/utility/update_utils"
	"testing"
)

var (
	baseUrl   = *InputBaseUrl
	githubTag = *InputGithubTag
	initData  = &g_structs.InitData{
		Department:   "GitHub",
		Name:         "Go-Test",
		LocalVersion: githubTag,
	}
)

func init() {
	// 传入后端地址
	g_consts.BaseUrl = baseUrl
}

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
		err := cli_utils.CmdCore.StartSpeedCmd(context.Background(), initData)
		t.Assert(err, nil)
	})
}
