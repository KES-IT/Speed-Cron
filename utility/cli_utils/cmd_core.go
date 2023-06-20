package cli_utils

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type uCmdCore struct{}

var CmdCore = &uCmdCore{}

// StartSpeedCmd
//
//	@dc:
//	@params:
//	@response:
//	@author: hamster   @date:2023/6/20 10:06:06
func (u *uCmdCore) StartSpeedCmd(ctx context.Context, initData g.Map) (err error) {
	cmd := CliUtils.CreateSpeedCmd()
	if cmd == nil {
		glog.Warning(ctx, "创建命令失败,获取测速节点失败")
		err = gerror.New("创建命令失败,获取测速节点失败")
		return
	}
	// 获取命令的标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取标准输出管道时发生错误:", err)
		return
	}
	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("启动命令时发生错误:", err)
		return
	}
	var (
		scanner             = bufio.NewScanner(stdout)
		defaultBars         = Bar.InitDefaultBar()
		uploadNetDataStruct = &NetInfoUploadData{}
	)
	uploadNetDataStruct.Department = gconv.String(initData["department"])
	uploadNetDataStruct.StaffName = gconv.String(initData["name"])
	// 持续获取输出
	for scanner.Scan() {
		// 获取输出行
		line := scanner.Bytes()
		ok, err := CmdProgress.CmdCoreProgress(ctx, string(line), defaultBars, uploadNetDataStruct)
		if err != nil {
			glog.Warning(ctx, "处理命令行输出时发生错误:", err)
			return err
		}
		if ok {
			return nil
		}
	}
	// 等待命令执行完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("等待命令执行完成时发生错误:", err)
		time.Sleep(time.Second * 5)
	}
	return
}
