// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TThirdDao is the data access object for table t_third.
type TThirdDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns TThirdColumns // columns contains all the column names of Table for convenient usage.
}

// TThirdColumns defines and stores column names for table t_third.
type TThirdColumns struct {
	Id       string //
	Name     string // 名称
	Status   string // 状态
	Key      string // key
	Url      string // 地址
	Merchant string // 商户号
	Token    string // 钱包验证token
	Remark   string //
}

// tThirdColumns holds the columns for table t_third.
var tThirdColumns = TThirdColumns{
	Id:       "id",
	Name:     "name",
	Status:   "status",
	Key:      "key",
	Url:      "url",
	Merchant: "merchant",
	Token:    "token",
	Remark:   "remark",
}

// NewTThirdDao creates and returns a new DAO object for table data access.
func NewTThirdDao() *TThirdDao {
	return &TThirdDao{
		group:   "default",
		table:   "t_third",
		columns: tThirdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TThirdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TThirdDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TThirdDao) Columns() TThirdColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TThirdDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TThirdDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TThirdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
