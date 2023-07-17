package boot

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
	"kes-cron/internal/global/g_cache"
	"kes-cron/internal/global/g_structs"
	"kes-cron/utility/cli_utils"
	"kes-cron/utility/cron_utils"
	"kes-cron/utility/net_utils"
	"kes-cron/utility/update_utils"
)

func Boot(initData *g_structs.InitData) (err error) {
	_, err = gcron.AddOnce(context.TODO(), "@every 1s", func(ctx context.Context) {
		glog.Debug(context.Background(), "定时任务启动中...")
		if err := bootMethod(initData); err != nil {
			glog.Fatal(context.Background(), "定时任务启动失败: ", err)
		}
		glog.Debug(context.Background(), "定时任务启动成功")
	}, "开始启动定时任务")
	if err != nil {
		return err
	}

	_, err = gcron.AddOnce(context.TODO(), "@every 3s", func(ctx context.Context) {
		glog.Info(context.Background(), "定时任务测试中...")
		if err := bootCheck(initData); err != nil {
			glog.Fatal(context.Background(), "定时任务测试失败: ", err)
		}
		glog.Info(context.Background(), "定时任务测试成功")
	}, "开始测试定时任务")
	if err != nil {
		return err
	}

	return nil
}

// bootCheck 测试初次启动任务
func bootCheck(initData *g_structs.InitData) (err error) {
	var ctx = context.Background()
	// 判断是否在更新中
	if gcache.MustGet(ctx, g_cache.UpdateCacheKey).Bool() {
		glog.Warning(ctx, "正在更新客户端程序，跳过本次测速")
		return
	}
	// 设置测速状态
	_ = gcache.Set(ctx, g_cache.SpeedCacheKey, true, 0)

	err = cli_utils.CmdCore.StartSpeedCmd(context.Background(), initData)
	if err != nil {
		glog.Error(context.Background(), "测试测速服务", err)
		return
	}
	// 移除测速状态
	_, _ = gcache.Remove(ctx, g_cache.SpeedCacheKey)

	err = net_utils.NetUtils.CoreLatency(initData)
	if err != nil {
		glog.Error(context.Background(), "HTTPS延迟检测失败: ", err)
		return
	}

	return nil
}

// bootMethod 初始化定时任务
func bootMethod(initData *g_structs.InitData) (err error) {
	var ctx = context.TODO()

	glog.Debug(ctx, "开始初始化定时任务管理器")
	_, err = gcron.AddSingleton(ctx, "@every 30s", func(ctx context.Context) {
		err := cron_utils.CronManage.GetConfigAndStart(ctx, initData)
		if err != nil {
			glog.Error(ctx, "初始化定时任务管理器服务失败: ", err)
			return
		}
	}, "Cron-Manager")
	if err != nil {
		glog.Warning(ctx, "添加初始化定时任务管理器服务失败: ", err)
		return err
	}
	glog.Debug(ctx, "初始化定时任务管理器服务成功")

	glog.Debug(ctx, "开始初始化自动更新服务")
	_, err = gcron.AddSingleton(ctx, "@every 5s", func(ctx context.Context) {
		err := update_utils.AutoUpdate.UpdateCore(ctx, initData)
		if err != nil {
			glog.Error(ctx, "自动更新服务失败: ", err)
			return
		}
		if !gcache.MustGet(ctx, g_cache.UpdateCacheKey).IsNil() {
			_, _ = gcache.Remove(ctx, g_cache.UpdateCacheKey)
		}
	}, "Cron-Update")
	if err != nil {
		glog.Warning(ctx, "添加初始化自动更新服务失败: ", err)
		return err
	}
	glog.Debug(ctx, "初始化自动更新服务成功")

	return nil
}
