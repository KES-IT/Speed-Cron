package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadSpeedDataReq 上传网速数据 Req请求
type UploadSpeedDataReq struct {
	g.Meta        `method:"post" tags:"网速数据" summary:"上传网速数据" dc:"上传网速数据"`
	MacAddress    string      `json:"mac_address"    v:"required #请输入 mac_address"` // mac地址
	DownloadSpeed float64     `json:"download_speed" `                              // 下载速度
	UploadSpeed   float64     `json:"upload_speed"   `                              // 上传速度
	Latency       string      `json:"latency"        `                              // 延迟
	Jitter        string      `json:"jitter"         `                              // 网络波动
	SpeedTime     *gtime.Time `json:"speed_time"     `                              // 测速时间
	ErrMsg        string      `json:"err_msg"        `                              // 错误信息
	TestDuration  string      `json:"test_duration"  `                              // 测速持续时间
	InternalIp    string      `json:"internal_ip"    `                              // 内网地址
	ExternalIp    string      `json:"external_ip"    `                              // 外网地址
}

// UploadSpeedDataRes 上传网速数据 Res返回
type UploadSpeedDataRes struct {
}
