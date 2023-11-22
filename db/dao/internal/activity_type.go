// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityTypeDao is the data access object for table a_activity_type.
type ActivityTypeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ActivityTypeColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityTypeColumns defines and stores column names for table a_activity_type.
type ActivityTypeColumns struct {
	Type   string //
	Name   string //
	Status string //
	Rule   string // 规则
	Remark string // 说明
}

// activityTypeColumns holds the columns for table a_activity_type.
var activityTypeColumns = ActivityTypeColumns{
	Type:   "type",
	Name:   "name",
	Status: "status",
	Rule:   "rule",
	Remark: "remark",
}

// NewActivityTypeDao creates and returns a new DAO object for table data access.
func NewActivityTypeDao() *ActivityTypeDao {
	return &ActivityTypeDao{
		group:   "default",
		table:   "a_activity_type",
		columns: activityTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityTypeDao) Columns() ActivityTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
