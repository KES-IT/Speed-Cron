// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NetSpeedDataDao is the data access object for table net_speed_data.
type NetSpeedDataDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns NetSpeedDataColumns // columns contains all the column names of Table for convenient usage.
}

// NetSpeedDataColumns defines and stores column names for table net_speed_data.
type NetSpeedDataColumns struct {
	Id            string // 网速表主键
	MacAddress    string // mac地址
	DownloadSpeed string // 下载速度
	UploadSpeed   string // 上传速度
	Latency       string // 延迟
	Jitter        string // 网络波动
	SpeedTime     string // 测速时间
	CreateTime    string // 创建时间
	ErrMsg        string // 错误信息
	TestDuration  string // 测速持续时间
	InternalIp    string // 内网地址
	ExternalIp    string // 外网地址
}

// netSpeedDataColumns holds the columns for table net_speed_data.
var netSpeedDataColumns = NetSpeedDataColumns{
	Id:            "id",
	MacAddress:    "mac_address",
	DownloadSpeed: "download_speed",
	UploadSpeed:   "upload_speed",
	Latency:       "latency",
	Jitter:        "jitter",
	SpeedTime:     "speed_time",
	CreateTime:    "create_time",
	ErrMsg:        "err_msg",
	TestDuration:  "test_duration",
	InternalIp:    "internal_ip",
	ExternalIp:    "external_ip",
}

// NewNetSpeedDataDao creates and returns a new DAO object for table data access.
func NewNetSpeedDataDao() *NetSpeedDataDao {
	return &NetSpeedDataDao{
		group:   "default",
		table:   "net_speed_data",
		columns: netSpeedDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NetSpeedDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NetSpeedDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NetSpeedDataDao) Columns() NetSpeedDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NetSpeedDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NetSpeedDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NetSpeedDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
