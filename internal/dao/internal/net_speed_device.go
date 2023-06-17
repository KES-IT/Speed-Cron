// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NetSpeedDeviceDao is the data access object for table net_speed_device.
type NetSpeedDeviceDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns NetSpeedDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// NetSpeedDeviceColumns defines and stores column names for table net_speed_device.
type NetSpeedDeviceColumns struct {
	Id             string // 设备id
	NetStatus      string // 在线状态
	MacAddress     string // 设备mac地址
	InternalIp     string // 设备内网ip
	OfflineTime    string // 离线时间
	LastOnlineTime string // 上次在线时间
}

// netSpeedDeviceColumns holds the columns for table net_speed_device.
var netSpeedDeviceColumns = NetSpeedDeviceColumns{
	Id:             "id",
	NetStatus:      "net_status",
	MacAddress:     "mac_address",
	InternalIp:     "internal_ip",
	OfflineTime:    "offline_time",
	LastOnlineTime: "last_online_time",
}

// NewNetSpeedDeviceDao creates and returns a new DAO object for table data access.
func NewNetSpeedDeviceDao() *NetSpeedDeviceDao {
	return &NetSpeedDeviceDao{
		group:   "default",
		table:   "net_speed_device",
		columns: netSpeedDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NetSpeedDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NetSpeedDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NetSpeedDeviceDao) Columns() NetSpeedDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NetSpeedDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NetSpeedDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NetSpeedDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
