package cli_utils

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"os"
)

type uBar struct{}

var Bar = &uBar{}

type DefaultBar struct {
	PingBar               *progressbar.ProgressBar
	DownloadBar           *progressbar.ProgressBar
	UploadBar             *progressbar.ProgressBar
	UploadNetSpeedDataBar *progressbar.ProgressBar
}

func (b *uBar) InitDefaultBar() *DefaultBar {
	return &DefaultBar{
		PingBar:               b.CreatePingBar(),
		DownloadBar:           b.CreateDownloadBar(),
		UploadBar:             b.CreateUploadBar(),
		UploadNetSpeedDataBar: b.CreateUploadNetSpeedDataBar(),
	}
}

func (b *uBar) CreatePingBar() *progressbar.ProgressBar {
	return progressbar.NewOptions(100,
		progressbar.OptionSetWriter(os.Stdout),
		progressbar.OptionEnableColorCodes(false),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetWidth(14),
		progressbar.OptionOnCompletion(func() {
			_, _ = fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSetDescription("[speed][1/4][ping] 测试延迟中..."))
}

func (b *uBar) CreateDownloadBar() *progressbar.ProgressBar {
	return progressbar.NewOptions(100,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionOnCompletion(func() {
			_, _ = fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSetDescription("[speed][2/4][download] 下行速度测试中..."))
}

func (b *uBar) CreateUploadBar() *progressbar.ProgressBar {
	return progressbar.NewOptions(100,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionOnCompletion(func() {
			_, _ = fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSetDescription("[speed][3/4][upload] 上行速度测试中..."))
}

func (b *uBar) CreateUploadNetSpeedDataBar() *progressbar.ProgressBar {
	return progressbar.NewOptions(100,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionOnCompletion(func() {
			_, _ = fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSetDescription("[speed][4/4][confirm] 上传数据中..."))
}
