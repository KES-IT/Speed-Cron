package s_speed_data_service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"kes-network-watcher/internal/dao"
	"kes-network-watcher/internal/global/g_functions"
	"kes-network-watcher/internal/model/entity"
	"kes-network-watcher/internal/service"
)

// ==========================================================================
// logic 初始化
// ==========================================================================

type sNetSpeedData struct {
}

func init() {
	service.RegisterNetSpeedData(New())
}

func New() service.INetSpeedData {
	return &sNetSpeedData{}
}

// AddNetSpeedData
//
//	@dc: 新增网络数据
//	@params:
//	@response:
//	@author:Administrator @date:2023-06-15 17:08:17
func (s *sNetSpeedData) AddNetSpeedData(ctx context.Context, in *entity.NetSpeedData) (out string, err error) {
	db := dao.NetSpeedData.Ctx(ctx)
	if err != nil {
		return "", err
	}
	id, err := db.OmitEmpty().InsertAndGetId(in)
	if err != nil {
		return "", fmt.Errorf("sCanteenFood AddAndUpdate 数据库新增错误 %v", err)
	} else {
		out = gconv.String(id)
	}

	return
}

// SelectById
//
//	@dc:主键查询表信息
//	@params:表主键id-in
//	@response:表原结构信息entity
//	@author:auto @date:2023-06-15 17:03:06
func (s *sNetSpeedData) SelectById(ctx context.Context, in *string) (out *entity.NetSpeedData, err error) {
	db := dao.NetSpeedData.Ctx(ctx)
	err = db.Where("id", in).Scan(&out)
	if err != nil {
		err = g_functions.ResErr(500, "sNetSpeedData SelectById 数据库查询错误！", err)
		return
	}
	return
}
