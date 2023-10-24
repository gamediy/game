// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WalletChangeLogDao is the data access object for table w_wallet_change_log.
type WalletChangeLogDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns WalletChangeLogColumns // columns contains all the column names of Table for convenient usage.
}

// WalletChangeLogColumns defines and stores column names for table w_wallet_change_log.
type WalletChangeLogColumns struct {
	Id              string //
	TransId         string // 订单号，（有可能是充值提现的订单号）
	Uid             string // UID
	Account         string // 账号
	Pid             string // 上级ID
	TransCode       string // 余额编号
	Title           string // 标题
	BalanceBefore   string // 之前余额
	BalanceAfter    string // 之后余额
	Note            string // 说明
	OrderNoRelation string // 关联订单号，（有可能是充值提现的订单号）
	ParentPath      string // 上级全路经
	Amount          string // 金额
	CreatedAt       string // 创建时间
	UpdateAt        string //
}

// walletChangeLogColumns holds the columns for table w_wallet_change_log.
var walletChangeLogColumns = WalletChangeLogColumns{
	Id:              "id",
	TransId:         "trans_id",
	Uid:             "uid",
	Account:         "account",
	Pid:             "pid",
	TransCode:       "trans_code",
	Title:           "title",
	BalanceBefore:   "balance_before",
	BalanceAfter:    "balance_after",
	Note:            "note",
	OrderNoRelation: "order_no_relation",
	ParentPath:      "parent_path",
	Amount:          "amount",
	CreatedAt:       "created_at",
	UpdateAt:        "update_at",
}

// NewWalletChangeLogDao creates and returns a new DAO object for table data access.
func NewWalletChangeLogDao() *WalletChangeLogDao {
	return &WalletChangeLogDao{
		group:   "default",
		table:   "w_wallet_change_log",
		columns: walletChangeLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WalletChangeLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WalletChangeLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WalletChangeLogDao) Columns() WalletChangeLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WalletChangeLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WalletChangeLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WalletChangeLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
