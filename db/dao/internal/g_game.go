// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GGameDao is the data access object for table g_game.
type GGameDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns GGameColumns // columns contains all the column names of Table for convenient usage.
}

// GGameColumns defines and stores column names for table g_game.
type GGameColumns struct {
	Code            string // 游戏编号
	Name            string // 名称
	Status          string // 状态0：关闭，1开始
	CreatedAt       string // 创建时间
	StartTime       string // 开始时间
	EndTime         string // 结束时间
	TotalIssue      string // 一共有多少期
	IntervalSeconds string // 每期间隔秒数
	TypeCode        string // 类型code
	Sort            string // 排序
	CategoryCode    string // 分类
}

// gGameColumns holds the columns for table g_game.
var gGameColumns = GGameColumns{
	Code:            "code",
	Name:            "name",
	Status:          "status",
	CreatedAt:       "created_at",
	StartTime:       "start_time",
	EndTime:         "end_time",
	TotalIssue:      "total_issue",
	IntervalSeconds: "interval_seconds",
	TypeCode:        "type_code",
	Sort:            "sort",
	CategoryCode:    "category_code",
}

// NewGGameDao creates and returns a new DAO object for table data access.
func NewGGameDao() *GGameDao {
	return &GGameDao{
		group:   "default",
		table:   "g_game",
		columns: gGameColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GGameDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GGameDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GGameDao) Columns() GGameColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GGameDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GGameDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GGameDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
