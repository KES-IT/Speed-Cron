package g_test

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
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
	initData = &g_structs.InitData{
		Department:   "GitHub",
		Name:         "Go-Test",
		LocalVersion: "",
	}
)

func init() {
	baseUrl := *InputBaseUrl
	githubTag := *InputGithubTag
	g.Dump(baseUrl)
	g.Dump(githubTag)
	g.Dump(initData)
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
		// 传入后端地址
		baseUrl := *InputBaseUrl
		g_consts.BaseUrl = baseUrl
		err := cli_utils.CmdCore.StartSpeedCmd(context.Background(), initData)
		t.Assert(err, nil)
	})
}

// 测试多站点延迟测试
func Test_Website_Latency(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 传入后端地址
		baseUrl := *InputBaseUrl
		g_consts.BaseUrl = baseUrl
		err := net_utils.NetUtils.CoreLatency(context.Background(), initData)
		t.Assert(err, nil)
	})
}

// 测试自动更新模块
func Test_Auto_Update(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 传入后端地址
		baseUrl := *InputBaseUrl
		g_consts.BaseUrl = baseUrl
		githubTag := *InputGithubTag
		initData.LocalVersion = githubTag
		_ = gcache.Set(context.Background(), "GitHubTestStatus", true, 0)
		err := update_utils.AutoUpdate.UpdateCore(context.Background(), initData)
		t.Assert(err, nil)
	})
}
