// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GSlotOrderDao is the data access object for table g_slot_order.
type GSlotOrderDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns GSlotOrderColumns // columns contains all the column names of Table for convenient usage.
}

// GSlotOrderColumns defines and stores column names for table g_slot_order.
type GSlotOrderColumns struct {
	OrderNo    string // 编号
	Uid        string // 用户编号
	Account    string // 用户账号
	CreatedAt  string // 时间
	Amount     string // 投注金额
	Profit     string // 赢利
	Status     string // 状态 0 撤单，1:投注成功，2：中奖，3：未中奖
	Odds       string // 赔率
	Fee        string // 手续费
	BetResult  string // 下注结果
	PlayCode   string // PlayId
	Pid        string //
	ParentPath string //
}

// gSlotOrderColumns holds the columns for table g_slot_order.
var gSlotOrderColumns = GSlotOrderColumns{
	OrderNo:    "order_no",
	Uid:        "uid",
	Account:    "account",
	CreatedAt:  "created_at",
	Amount:     "amount",
	Profit:     "profit",
	Status:     "status",
	Odds:       "odds",
	Fee:        "fee",
	BetResult:  "bet_result",
	PlayCode:   "play_code",
	Pid:        "pid",
	ParentPath: "parent_path",
}

// NewGSlotOrderDao creates and returns a new DAO object for table data access.
func NewGSlotOrderDao() *GSlotOrderDao {
	return &GSlotOrderDao{
		group:   "default",
		table:   "g_slot_order",
		columns: gSlotOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GSlotOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GSlotOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GSlotOrderDao) Columns() GSlotOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GSlotOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GSlotOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GSlotOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
