package g_consts

import "github.com/gogf/gf/v2/frame/g"

var (
	JWTKey = []byte("hamster-im")
)

type DefaultActionMessage struct {
	Message string `json:"message" dc:"返回信息"`
}

const (
	ByteToG = 1024 * 1024 * 1024
)

const (
	PushCoreUrl = "http://120.24.211.49:10399/PushCore"
	// LocalPushCoreUrl = "http://127.0.0.1:10399/PushCore"
)

// home_network.go配置文件
const (
	HomeNetworkRouterIP       = "router.buycoffee.top:580"
	HomeNetworkRouterAddress  = "http://router.buycoffee.top:580"
	HomeNetworkRouterPassword = "deny1963"
)

// node_utils.go配置文件
const (
	XrayBaseUrl      = "ray.buycoffee.top:580"
	XrayLoginDataMap = `{"username":"hamster","password":"deny1963"}`
)

// proxy_network.go配置文件
var (
	XuiBaseUrl      = "http://xui.buycoffee.top:580"
	XuiLoginDataMap = g.Map{
		"username": "hamster",
		"password": "Deny1963!",
	}
)

// coffee.go配置文件
const (
	CoffeeBaseUrl  = "https://portal.coffeecloud.top/api/v1/user/getSubscribe"
	CoffeeLoginUrl = "https://portal.coffeecloud.top/api/v1/passport/auth/login"
)

var (
	CoffeeAuthData = g.Map{
		"email":    "liaolaixin@gmail.com",
		"password": "deny1963",
	}
)

// cache key
const (
	HomeNetworkCacheKey   = "homeNetwork"
	ProxyNetworkCacheKey  = "proxyNetwork"
	ProxySessionCacheKey  = "proxySession"
	ProxyCountCacheKey    = "proxyNetworkUpSpeedCount"
	ProxyUserFlowCacheKey = "proxyUserFlow"
	ProxyOutboundCacheKey = "proxyOutbound"
	ProxyNodeCacheKey     = "proxyNode"
	CoffeeCacheKey        = "coffee"
)
