// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GGameCategoryDao is the data access object for table g_game_category.
type GGameCategoryDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns GGameCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// GGameCategoryColumns defines and stores column names for table g_game_category.
type GGameCategoryColumns struct {
	Code   string //
	Status string //
	Logo   string //
	Name   string //
}

// gGameCategoryColumns holds the columns for table g_game_category.
var gGameCategoryColumns = GGameCategoryColumns{
	Code:   "code",
	Status: "status",
	Logo:   "logo",
	Name:   "name",
}

// NewGGameCategoryDao creates and returns a new DAO object for table data access.
func NewGGameCategoryDao() *GGameCategoryDao {
	return &GGameCategoryDao{
		group:   "default",
		table:   "g_game_category",
		columns: gGameCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GGameCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GGameCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GGameCategoryDao) Columns() GGameCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GGameCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GGameCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GGameCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
