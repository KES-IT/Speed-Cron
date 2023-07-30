package cli_utils

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"os"
)

type uBar struct{}

var Bar = &uBar{}

// DefaultBar 默认进度条
type DefaultBar struct {
	PingBar               *progressbar.ProgressBar
	DownloadBar           *progressbar.ProgressBar
	UploadBar             *progressbar.ProgressBar
	UploadNetSpeedDataBar *progressbar.ProgressBar
}

// InitDefaultBar 初始化默认进度条
func (b *uBar) InitDefaultBar() *DefaultBar {
	return &DefaultBar{
		PingBar:               b.CreatePingBar(),
		DownloadBar:           b.CreateDownloadBar(),
		UploadBar:             b.CreateUploadBar(),
		UploadNetSpeedDataBar: b.CreateUploadNetSpeedDataBar(),
	}
}

// CreatePingBar 创建ping进度条
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

// CreateDownloadBar 创建下载进度条
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

// CreateUploadBar 创建上传进度条
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

// CreateUploadNetSpeedDataBar 创建上传网速数据进度条
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
