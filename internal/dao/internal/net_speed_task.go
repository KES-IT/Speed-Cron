// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NetSpeedTaskDao is the data access object for table net_speed_task.
type NetSpeedTaskDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns NetSpeedTaskColumns // columns contains all the column names of Table for convenient usage.
}

// NetSpeedTaskColumns defines and stores column names for table net_speed_task.
type NetSpeedTaskColumns struct {
	Id               string // 任务列表id
	Type             string // 任务类型
	Command          string // 任务指令
	Device           string // 设备id
	DeviceMacAddress string // 设备mac地址
	TaskStatus       string // 任务状态
	PullTime         string // 拉取时间
	FinishedTime     string // 完成时间
	CreateTime       string // 创建时间
}

// netSpeedTaskColumns holds the columns for table net_speed_task.
var netSpeedTaskColumns = NetSpeedTaskColumns{
	Id:               "id",
	Type:             "type",
	Command:          "command",
	Device:           "device",
	DeviceMacAddress: "device_mac_address",
	TaskStatus:       "task_status",
	PullTime:         "pull_time",
	FinishedTime:     "finished_time",
	CreateTime:       "create_time",
}

// NewNetSpeedTaskDao creates and returns a new DAO object for table data access.
func NewNetSpeedTaskDao() *NetSpeedTaskDao {
	return &NetSpeedTaskDao{
		group:   "default",
		table:   "net_speed_task",
		columns: netSpeedTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NetSpeedTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NetSpeedTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NetSpeedTaskDao) Columns() NetSpeedTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NetSpeedTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NetSpeedTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NetSpeedTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
