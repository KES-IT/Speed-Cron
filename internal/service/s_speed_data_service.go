// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"kes-network-watcher/internal/model/entity"
)

type (
	INetSpeedData interface {
		AddNetSpeedData(ctx context.Context, in *entity.NetSpeedData) (out string, err error)
		SelectById(ctx context.Context, in *string) (out *entity.NetSpeedData, err error)
	}
)

var (
	localNetSpeedData INetSpeedData
)

func NetSpeedData() INetSpeedData {
	if localNetSpeedData == nil {
		panic("implement not found for interface INetSpeedData, forgot register?")
	}
	return localNetSpeedData
}

func RegisterNetSpeedData(i INetSpeedData) {
	localNetSpeedData = i
}
