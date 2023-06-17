package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
	"kes-cron/utility/cli_utils"
	"kes-cron/utility/net_utils"
)

func Boot(initData g.Map) (err error) {
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

func bootCheck(initData g.Map) (err error) {
	err = cli_utils.CliUtils.StartSingleSpeedTest(initData)
	if err != nil {
		glog.Error(context.Background(), "测试测速服务", err)
		return
	}
	return nil
}

func bootMethod(initData g.Map) (err error) {
	var ctx = context.TODO()

	glog.Notice(ctx, "开始HTTPS延迟定时检测服务")
	_, err = gcron.AddSingleton(ctx, "@every 15s", func(ctx context.Context) {
		err := net_utils.NetUtils.CoreLatency(initData)
		if err != nil {
			glog.Error(ctx, "HTTPS延迟检测失败: ", err)
			return
		}
	}, "HTTPS延迟检测服务")
	if err != nil {
		return err
	}

	glog.Notice(ctx, "开始定时(1h)测速服务")
	_, err = gcron.AddSingleton(ctx, "@every 1h", func(ctx context.Context) {
		err := cli_utils.CliUtils.StartSingleSpeedTest(initData)
		if err != nil {
			glog.Error(ctx, "定时测速服务失败: ", err)
			return
		}
	}, "定时测速服务")
	if err != nil {
		return err
	}

	return nil
}
