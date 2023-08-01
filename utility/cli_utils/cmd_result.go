package cli_utils

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-cron/internal/global/g_consts"
	"kes-cron/utility/net_utils"
)

type uResult struct{}

var Result = &uResult{}

type NetInfoUploadData struct {
	PingJitter        string `json:"jitter" description:"ping抖动"`
	PingLatency       string `json:"latency" description:"ping延迟"`
	DownloadBandwidth string `json:"download_speed" description:"下载带宽"`
	UploadBandwidth   string `json:"upload_speed" description:"上传带宽"`
	MacAddress        string `json:"mac_address" description:"mac地址"`
	ErrMsg            string `json:"err_msg" description:"错误信息"`
	TimeStamp         string `json:"speed_time" description:"时间戳"`
	InternalIp        string `json:"internal_ip" description:"内网IP"`
	ExternalIp        string `json:"external_ip" description:"外网IP"`
	Department        string `json:"department" description:"部门"`
	StaffName         string `json:"staff_name" description:"姓名"`
}

// ProcessResult
//
//	@dc: 处理结果
//	@author: hamster   @date:2023-06-16 09:45:06
func (u *uResult) ProcessResult(oriData *NetInfoUploadData) (err error) {
	// 进行数据处理
	if oriData.InternalIp == "" || oriData.MacAddress == "00:00:00:00:00:00" {
		oriData.InternalIp, oriData.MacAddress = net_utils.NetworkInfo.GetMacAddress()
	}
	if oriData.Department == "" {
		oriData.Department = "未知"
	}
	if oriData.StaffName == "" {
		oriData.StaffName = "未知"
	}
	oriData.TimeStamp = gtime.Now().String()
	glog.Info(context.Background(), "测试完成")
	glog.Info(context.Background(), "网络延迟为：", oriData.PingLatency, "ms")
	glog.Info(context.Background(), "网络抖动为：", oriData.PingJitter, "ms")
	glog.Info(context.Background(), "下载速度为：", oriData.DownloadBandwidth, "Mbps")
	glog.Info(context.Background(), "上传速度为：", oriData.UploadBandwidth, "Mbps")
	glog.Info(context.Background(), "内网IP地址为：", oriData.InternalIp, "，MAC地址为：", oriData.MacAddress, "，外网IP地址为：", oriData.ExternalIp)

	// 2.上传数据到服务器
	err = u.UploadDataToServer(oriData)
	if err != nil {
		glog.Warning(context.Background(), err)
		return
	}
	return
}

// UploadDataToServer 上传数据到服务器
func (u *uResult) UploadDataToServer(netInfo *NetInfoUploadData) error {
	// 构建进度条
	dataBar := Bar.CreateUploadNetSpeedDataBar()
	// 转换为json
	uploadJson := gjson.New(netInfo).MapStrAny()
	glog.Debug(context.Background(), "---上传数据---")
	g.Dump(uploadJson)
	glog.Debug(context.Background(), "---上传数据---")
	// 上传数据
	post, err := g.Client().Post(context.Background(), g_consts.BackendBaseUrl()+g_consts.SpeedBackendUrl, uploadJson)
	if err != nil {
		return err
	}
	defer func(post *gclient.Response) {
		err := post.Close()
		if err != nil {
			glog.Warning(context.Background(), "上传测速数据关闭连接失败", err)
		}
	}(post)
	if post.StatusCode != 200 {
		return errors.New("上传数据失败,状态码：" + gconv.String(post.StatusCode))
	}
	// 上传成功，修改进度条
	_ = dataBar.Finish()
	glog.Info(context.Background(), "上传数据成功")
	return nil
}
