// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLoginLogDao is the data access object for table u_user_login_log.
type UserLoginLogDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UserLoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// UserLoginLogColumns defines and stores column names for table u_user_login_log.
type UserLoginLogColumns struct {
	Id        string //
	Uid       string //
	Account   string //
	Ip        string //
	CreatedAt string //
}

// userLoginLogColumns holds the columns for table u_user_login_log.
var userLoginLogColumns = UserLoginLogColumns{
	Id:        "id",
	Uid:       "uid",
	Account:   "account",
	Ip:        "ip",
	CreatedAt: "created_at",
}

// NewUserLoginLogDao creates and returns a new DAO object for table data access.
func NewUserLoginLogDao() *UserLoginLogDao {
	return &UserLoginLogDao{
		group:   "default",
		table:   "u_user_login_log",
		columns: userLoginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserLoginLogDao) Columns() UserLoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
