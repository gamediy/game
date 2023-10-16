// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameTypeDao is the data access object for table g_game_type.
type GameTypeDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns GameTypeColumns // columns contains all the column names of Table for convenient usage.
}

// GameTypeColumns defines and stores column names for table g_game_type.
type GameTypeColumns struct {
	Code   string //
	Status string //
	Logo   string //
	Name   string //
}

// gameTypeColumns holds the columns for table g_game_type.
var gameTypeColumns = GameTypeColumns{
	Code:   "code",
	Status: "status",
	Logo:   "logo",
	Name:   "name",
}

// NewGameTypeDao creates and returns a new DAO object for table data access.
func NewGameTypeDao() *GameTypeDao {
	return &GameTypeDao{
		group:   "default",
		table:   "g_game_type",
		columns: gameTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameTypeDao) Columns() GameTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
