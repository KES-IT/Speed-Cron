package boot

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
)

func Boot() (err error) {
	_, err = gcron.AddOnce(context.TODO(), "@every 1s", func(ctx context.Context) {
		glog.Debug(context.Background(), "定时任务启动中...")
		if err := bootMethod(); err != nil {
			glog.Fatal(context.Background(), "定时任务启动失败: ", err)
		}
		glog.Debug(context.Background(), "定时任务启动成功")
	}, "开始启动定时任务")
	if err != nil {
		return err
	}

	_, err = gcron.AddOnce(context.TODO(), "@every 10s", func(ctx context.Context) {
		glog.Info(context.Background(), "定时任务测试中...")
		if err := bootCheck(); err != nil {
			glog.Fatal(context.Background(), "定时任务测试失败: ", err)
		}
		glog.Info(context.Background(), "定时任务测试成功")
	}, "开始测试定时任务")
	if err != nil {
		return err
	}

	return nil
}

func bootCheck() (err error) {
	return nil
}

func bootMethod() (err error) {
	return nil
}
