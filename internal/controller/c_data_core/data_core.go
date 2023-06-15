package c_data_core

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "kes-network-watcher/api/v1"
	"kes-network-watcher/internal/global/g_functions"
	"kes-network-watcher/internal/model/entity"
	"kes-network-watcher/internal/service"
)

type cDataCore struct{}

var DataCore = &cDataCore{}

// UploadSpeedData
//
//	@dc: 上传网速测试数据
//	@author: Administrator   @date:2023-06-15 16:51:26
func (c *cDataCore) UploadSpeedData(ctx context.Context, req *v1.UploadSpeedDataReq) (res *v1.UploadSpeedDataRes, err error) {
	inputData := entity.NetSpeedData{}
	err = gconv.Struct(req, &inputData)
	if err != nil {
		err = g_functions.ResErr(400, "cDataCore UploadSpeedData 数据转换错误！", err)
		return nil, err
	}
	_, err = service.NetSpeedData().AddNetSpeedData(ctx, &inputData)
	if err != nil {
		err = g_functions.ResErr(400, "cDataCore UploadSpeedData 数据库新增错误！", err)
		return nil, err
	}
	return
}
