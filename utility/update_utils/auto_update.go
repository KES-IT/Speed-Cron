package update_utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/kardianos/osext"
	"gopkg.in/inconshreveable/go-update.v0"
	"kes-cron/internal/global/g_cache"
	"kes-cron/internal/global/g_consts"
	"kes-cron/internal/global/g_structs"
	"kes-cron/utility/net_utils"
	"os"
	"path/filepath"
	"time"
)

type uAutoUpdate struct{}

var AutoUpdate = &uAutoUpdate{}

// UpdateCore
//
//	@dc: 更新任务核心程序
//	@author: laixin   @date:2023/7/14 09:37:24
func (u *uAutoUpdate) UpdateCore(ctx context.Context, initData *g_structs.InitData) (err error) {
	// 检测是否在进行测速服务
	if gcache.MustGet(ctx, g_cache.SpeedCacheKey).Bool() {
		glog.Warning(ctx, "正在进行测速服务，无法更新")
		return nil
	}

	// 获取设备信息
	_, macAddress := net_utils.NetworkInfo.GetMacAddress()
	// 获取配置
	glog.Debug(context.TODO(), "获取设备更新通道，当前mac地址为", macAddress)
	// 获取后端地址
	baseUrl := gcache.MustGet(context.Background(), "BackendBaseUrl").String()
	configResponse, err := g.Client().SetTimeout(5*time.Second).Post(context.TODO(), baseUrl+g_consts.ConfigBackendUrl, g.Map{
		"mac_address": macAddress,
	})
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			glog.Warning(context.TODO(), "UpdateCore getConfig 关闭response失败: ", err)
		}
	}(configResponse)
	if err != nil {
		return
	}
	if configResponse.StatusCode != 200 {
		err = gerror.Newf("UpdateCore getConfig 获取配置失败，错误码：%d", configResponse.StatusCode)
		return
	}
	// 解析配置
	configJson := gjson.New(configResponse.ReadAllString())
	// 获取更新通道
	updateChannel := configJson.Get("data.update_channel").String()
	// 判断更新通道是否为空
	if updateChannel == "" {
		updateChannel = "stable"
	}

	var (
		githubVersion     string
		githubDownloadUrl string
		downloadStatus    bool
	)

	switch updateChannel {
	case "stable":
		glog.Debug(ctx, "当前更新通道为稳定版")
		// 获取最新GitHub Release版本信息与下载地址
		githubVersion, githubDownloadUrl, downloadStatus = getLatestVersionInfo(false)
	case "beta":
		glog.Debug(ctx, "当前更新通道为测试版")
		// 获取最新GitHub Release版本信息与下载地址
		githubVersion, githubDownloadUrl, downloadStatus = getLatestVersionInfo(true)
	}

	glog.Info(ctx, "githubVersion: ", githubVersion, " githubDownloadUrl: ", githubDownloadUrl, " downloadStatus: ", downloadStatus)
	if githubVersion == "" || githubDownloadUrl == "" || !downloadStatus {
		glog.Warning(ctx, "获取github最新版本失败，无法比较版本")
		return nil
	}

	glog.Info(ctx, "目前本地localVersion为: ", initData.LocalVersion, "目前最新githubVersion为: ", githubVersion)
	if githubVersion == initData.LocalVersion {
		glog.Info(ctx, "speed_cron版本是最新，无需下载...")
		return nil
	}

	// 设置更新状态缓存
	_ = gcache.Set(ctx, g_cache.UpdateCacheKey, true, 0)

	glog.Debug(ctx, "开始更新speed_cron...")

	proxyDownloadUrl := g_consts.DownloadProxyUrl + githubDownloadUrl
	glog.Debug(ctx, "proxyDownloadUrl: ", proxyDownloadUrl)

	err = updateFunc(proxyDownloadUrl)
	if err != nil {
		glog.Warning(ctx, "更新speed_cron失败，原因：", err.Error())
		return
	}
	return
}

// updateFunc 更新speed_cron二进制程序
func updateFunc(downloadUrl string) error {
	// 获取当前程序路径
	path, err := osext.Executable()
	if err != nil {
		return err
	}
	if resolvedPath, err := filepath.EvalSymlinks(path); err == nil {
		path = resolvedPath
	}

	old, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(old *os.File) {
		err := old.Close()
		if err != nil {
			glog.Warning(context.TODO(), "关闭旧文件失败，原因：", err.Error())
		}
	}(old)

	// 下载最新的speed_cron
	exe, err := g.Client().Get(context.TODO(), downloadUrl)
	if err != nil {
		return err
	}
	if exe.StatusCode != 200 {
		return fmt.Errorf("download failed: %s", exe.Status)
	}
	if exe.ContentLength < 1024 {
		return fmt.Errorf("download failed, file too small")
	}
	bin := exe.ReadAll()
	if len(bin) != int(exe.ContentLength) {
		return fmt.Errorf("download failed, file size mismatch")
	}
	glog.Debug(context.TODO(), "下载最新的speed_cron成功！")

	// 在windows上需要关闭旧文件才能更新
	_ = old.Close()
	// 更新speed_cron
	err, errRecover := update.New().FromStream(bytes.NewBuffer(bin))
	if errRecover != nil {
		return fmt.Errorf("update and recovery errors: %q %q", err, errRecover)
	}
	if err != nil {
		return err
	}

	glog.Debug(context.TODO(), "更新完成,重启中......")

	// 采用os.Exit(1)方式退出，等待winsw接管重启
	time.Sleep(1 * time.Second)
	os.Exit(1)

	return nil
}

// getLatestVersion 获取github最新版本
func getLatestVersionInfo(isBeta bool) (version string, downloadUrl string, downloadStatus bool) {
	// 获取后端地址
	baseUrl := gcache.MustGet(context.Background(), "BackendBaseUrl").String()
	backendURL := baseUrl + g_consts.StableBackendUrl

	// 判断是否为测试版
	if isBeta {
		backendURL = baseUrl + g_consts.BetaBackendUrl
	}

	response, err := g.Client().Get(context.TODO(), backendURL)
	if err != nil {
		glog.Warning(context.TODO(), "请求github最新版本失败，原因：", err.Error())
		return "", "", false
	}
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			glog.Warning(context.TODO(), "关闭response失败，原因：", err.Error())
		}
	}(response)
	githubResJson, err := gjson.DecodeToJson(response.ReadAllString())
	if err != nil {
		glog.Warning(context.TODO(), "解析response失败，原因：", err.Error())
		return "", "", false
	}

	// 判断GitHub Release可更新二进制文件是否存在
	if len(githubResJson.Get("data.github_res.assets").Array()) == 0 {
		glog.Warning(context.TODO(), "解析response失败，原因：", "github_res.assets为空")
		return "", "", false
	}
	version = githubResJson.Get("data.github_res.tag_name").String()

	// 获取下载文件名是否正确
	downloadFileName := githubResJson.Get("data.github_res.assets.0.name").String()
	if downloadFileName != g_consts.DownloadFileName {
		glog.Warning(context.TODO(), "解析response失败，原因：", "downloadFileName不正确")
		return "", "", false
	}

	// 获取下载地址
	downloadUrl = githubResJson.Get("data.github_res.assets.0.browser_download_url").String()
	if version == "" || downloadUrl == "" {
		glog.Warning(context.TODO(), "解析response失败，原因：", "version或downloadUrl为空")
		return "", "", false
	}
	return version, downloadUrl, true
}
