package cli_utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

type uCmdProgress struct{}

var CmdProgress = &uCmdProgress{}

// CmdCoreProgress
//
//	@dc: 命令行返回处理
//	@author: hamster   @date:2023-06-16 13:50:57
func (u *uCmdProgress) CmdCoreProgress(ctx context.Context, cmdString string, progressBar *DefaultBar, netInfoStruct *NetInfoUploadData) (status bool, err error) {
	ok, err := u.CmdLineProgress(ctx, cmdString, progressBar, netInfoStruct)
	if err != nil {
		netInfoStruct.ErrMsg = err.Error()
		err := Result.UploadDataToServer(netInfoStruct)
		if err != nil {
			glog.Warning(context.Background(), "上报错误失败")
			return false, err
		}
		return false, err
	}
	if ok {
		return true, nil
	}
	return false, nil
}

// CmdLineProgress
//
//	@dc: 处理测速命令行输出数据
//	@author: hamster   @date:2023-06-16 13:49:41
func (u *uCmdProgress) CmdLineProgress(ctx context.Context, cmdString string, progressBar *DefaultBar, netInfoStruct *NetInfoUploadData) (status bool, err error) {
	outPutJson := gjson.New(cmdString)
	testType := outPutJson.Get("type").String()
	switch testType {
	case "testStart":
		netInfoStruct.InternalIp = outPutJson.Get("interface.internalIp").String()
		netInfoStruct.ExternalIp = outPutJson.Get("interface.externalIp").String()
		netInfoStruct.MacAddress = outPutJson.Get("interface.macAddr").String()
	case "ping":
		if gconv.Int(outPutJson.Get("ping.progress").Float64()*100) == 100 {
			netInfoStruct.PingJitter = outPutJson.Get("ping.jitter").String()
			netInfoStruct.PingLatency = outPutJson.Get("ping.latency").String()
		}
		_ = progressBar.PingBar.Set(gconv.Int(outPutJson.Get("ping.progress").Float64() * 100))
	case "download":
		if gconv.Int(outPutJson.Get("download.progress").Float64()*100) == 100 {
			netInfoStruct.DownloadBandwidth = gconv.String(outPutJson.Get("download.bandwidth").Int() / 1024 / 100)
		}
		_ = progressBar.DownloadBar.Set(gconv.Int(outPutJson.Get("download.progress").Float64() * 100))
	case "upload":
		if gconv.Int(outPutJson.Get("upload.progress").Float64()*100) >= 10 {
			netInfoStruct.UploadBandwidth = gconv.String(outPutJson.Get("upload.bandwidth").Int() / 1024 / 100)
			_ = progressBar.UploadBar.Finish()
			err = Result.ProcessResult(netInfoStruct)
			if err != nil {
				glog.Warning(ctx, err)
				return false, err
			}
			return true, nil
		}
		_ = progressBar.UploadBar.Set(gconv.Int(outPutJson.Get("upload.progress").Float64() * 100))

	}
	return false, nil
}
