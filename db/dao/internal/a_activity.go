// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AActivityDao is the data access object for table a_activity.
type AActivityDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AActivityColumns // columns contains all the column names of Table for convenient usage.
}

// AActivityColumns defines and stores column names for table a_activity.
type AActivityColumns struct {
	Id      string //
	Title   string //
	Type    string // 类型
	Content string //
	Images  string //
	Status  string //
	Event   string // 事件
	Rule    string //
}

// aActivityColumns holds the columns for table a_activity.
var aActivityColumns = AActivityColumns{
	Id:      "id",
	Title:   "title",
	Type:    "type",
	Content: "content",
	Images:  "images",
	Status:  "status",
	Event:   "event",
	Rule:    "rule",
}

// NewAActivityDao creates and returns a new DAO object for table data access.
func NewAActivityDao() *AActivityDao {
	return &AActivityDao{
		group:   "default",
		table:   "a_activity",
		columns: aActivityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AActivityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AActivityDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AActivityDao) Columns() AActivityColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AActivityDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AActivityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AActivityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
