// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NetSpeedData is the golang structure of table net_speed_data for DAO operations like Where/Data.
type NetSpeedData struct {
	g.Meta        `orm:"table:net_speed_data, do:true"`
	Id            interface{} // 网速表主键
	MacAddress    interface{} // mac地址
	DownloadSpeed interface{} // 下载速度
	UploadSpeed   interface{} // 上传速度
	Latency       interface{} // 延迟
	Jitter        interface{} // 网络波动
	SpeedTime     *gtime.Time // 测速时间
	CreateTime    *gtime.Time // 创建时间
	ErrMsg        interface{} // 错误信息
	TestDuration  interface{} // 测速持续时间
	InternalIp    interface{} // 内网地址
	ExternalIp    interface{} // 外网地址
}
