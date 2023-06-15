// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NetSpeedData is the golang structure for table net_speed_data.
type NetSpeedData struct {
	Id            int         `json:"id"             ` // 网速表主键
	MacAddress    string      `json:"mac_address"    ` // mac地址
	DownloadSpeed float64     `json:"download_speed" ` // 下载速度
	UploadSpeed   float64     `json:"upload_speed"   ` // 上传速度
	Latency       string      `json:"latency"        ` // 延迟
	Jitter        string      `json:"jitter"         ` // 网络波动
	SpeedTime     *gtime.Time `json:"speed_time"     ` // 测速时间
	CreateTime    *gtime.Time `json:"create_time"    ` // 创建时间
	ErrMsg        string      `json:"err_msg"        ` // 错误信息
	TestDuration  string      `json:"test_duration"  ` // 测速持续时间
	InternalIp    string      `json:"internal_ip"    ` // 内网地址
	ExternalIp    string      `json:"external_ip"    ` // 外网地址
}
