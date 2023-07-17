package update_utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/kardianos/osext"
	"gopkg.in/inconshreveable/go-update.v0"
	"kes-cron/internal/global/g_cache"
	"kes-cron/internal/global/g_consts"
	"kes-cron/internal/global/g_structs"
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

	// 本地版本与服务器版本比较
	githubVersion := getLatestVersion()
	if githubVersion == "" {
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
	err = updateFunc()
	if err != nil {
		glog.Warning(ctx, "更新speed_cron失败，原因：", err.Error())
		return
	}
	return
}

// updateFunc 更新speed_cron二进制程序
func updateFunc() error {
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
	exe, err := g.Client().Get(context.TODO(), g_consts.DownloadExeUrl)
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
	time.Sleep(5 * time.Second)
	os.Exit(1)

	return nil
}

// getLatestVersion 获取github最新版本
func getLatestVersion() (version string) {
	response, err := g.Client().Get(context.TODO(), g_consts.UpdateBackendUrl)
	if err != nil {
		glog.Warning(context.TODO(), "请求github最新版本失败，原因：", err.Error())
		return ""
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
		return ""
	}

	// 判断GitHub Release可更新二进制文件是否存在
	if len(githubResJson.Get("data.github_res.assets").Array()) == 0 {
		glog.Warning(context.TODO(), "解析response失败，原因：", "github_res.assets为空")
		return ""
	}

	return githubResJson.Get("data.github_res.tag_name").String()
}
