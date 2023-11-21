// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TransactionTypeDao is the data access object for table w_transaction_type.
type TransactionTypeDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns TransactionTypeColumns // columns contains all the column names of Table for convenient usage.
}

// TransactionTypeColumns defines and stores column names for table w_transaction_type.
type TransactionTypeColumns struct {
	Id        string //
	Title     string //
	Code      string //
	Remark    string //
	Type      string //
	Class     string //
	CreatedAt string //
	UpdatedAt string //
}

// transactionTypeColumns holds the columns for table w_transaction_type.
var transactionTypeColumns = TransactionTypeColumns{
	Id:        "id",
	Title:     "title",
	Code:      "code",
	Remark:    "remark",
	Type:      "implement",
	Class:     "class",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTransactionTypeDao creates and returns a new DAO object for table data access.
func NewTransactionTypeDao() *TransactionTypeDao {
	return &TransactionTypeDao{
		group:   "default",
		table:   "w_transaction_type",
		columns: transactionTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TransactionTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TransactionTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TransactionTypeDao) Columns() TransactionTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TransactionTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TransactionTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TransactionTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
