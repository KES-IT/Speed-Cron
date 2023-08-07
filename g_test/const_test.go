package g_test

import (
	"context"
	"flag"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"net/http"
	"testing"
)

var (
	inputGithubTag = flag.String("tag", "", "get github tag")
	inputBaseUrl   = flag.String("baseurl", "", "get backend base url")
)

// 测试是否正确获取到后端地址
func Test_BackendBaseUrl(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 获取命令行参数
		flag.Parse()
		// 测试是否正确获取到后端地址
		baseUrl := *inputBaseUrl
		g.Dump(baseUrl)
		t.AssertNE(baseUrl, "")
		// 测试是否正确获取到后端地址且可访问
		response, err := g.Client().Get(context.Background(), baseUrl)
		if err != nil {
			t.Error(err)
		}
		t.AssertEQ(response.StatusCode, http.StatusOK)
	})
}

// 测试是否正确获取版本号
func Test_Version(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		flag.Parse()
		githubTag := *inputGithubTag
		g.Dump(githubTag)
		t.AssertNE(githubTag, "")
	})
}
