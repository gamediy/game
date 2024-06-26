// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminLoginLogDao is the data access object for table s_admin_login_log.
type AdminLoginLogDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AdminLoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// AdminLoginLogColumns defines and stores column names for table s_admin_login_log.
type AdminLoginLogColumns struct {
	Id        string //
	Uid       string //
	Account   string //
	Ip        string //
	CreatedAt string //
	UpdatedAt string //
}

// adminLoginLogColumns holds the columns for table s_admin_login_log.
var adminLoginLogColumns = AdminLoginLogColumns{
	Id:        "id",
	Uid:       "uid",
	Account:   "account",
	Ip:        "ip",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminLoginLogDao creates and returns a new DAO object for table data access.
func NewAdminLoginLogDao() *AdminLoginLogDao {
	return &AdminLoginLogDao{
		group:   "default",
		table:   "s_admin_login_log",
		columns: adminLoginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminLoginLogDao) Columns() AdminLoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
