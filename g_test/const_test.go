package g_test

import (
	"context"
	"flag"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/test/gtest"
	"kes-cron/internal/global/g_consts"
	"kes-cron/internal/global/g_structs"
	"net/http"
	"testing"
)

var (
	InputGithubTagFlag = flag.String("tag", "", "get github tag")
	InputBaseUrlFlag   = flag.String("baseurl", "", "get backend base url")
)

var (
	InitData = &g_structs.InitData{
		Department:   "GitHub",
		Name:         "Go-Test",
		LocalVersion: "",
	}
	InputGithubTag = ""
	InputBaseUrl   = ""
)

func Test_Init(t *testing.T) {
	// 获取命令行参数
	flag.Parse()
	InputGithubTag = *InputGithubTagFlag
	InputBaseUrl = *InputBaseUrlFlag
	InitData.LocalVersion = InputGithubTag
	g_consts.BaseUrl = InputBaseUrl
	// 设置更新测试标识缓存
	_ = gcache.Set(context.Background(), "GitHubTestStatus", true, 0)
}

// 测试是否正确获取到后端地址
func Test_BackendBaseUrl(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 测试是否正确获取到后端地址
		t.AssertNE(InputBaseUrl, "")
		// 测试是否正确获取到后端地址且可访问
		response, err := g.Client().Get(context.Background(), InputBaseUrl)
		if err != nil {
			t.Error(err)
		}
		t.AssertEQ(response.StatusCode, http.StatusOK)
	})
}

// 测试是否正确获取版本号
func Test_Version(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(InputGithubTag, "")
	})
}
