package r_hamster_router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"kes-network-watcher/internal/controller/c_data_core"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		BindDataCore(group)
	})
}

// BindDataCore 注册核心数据路由
func BindDataCore(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		// 自定中间件设置
		// group.Middleware(middleware.JWTAuth)
		// Bind注册路由
		group.Bind(c_data_core.DataCore)
	})
}
