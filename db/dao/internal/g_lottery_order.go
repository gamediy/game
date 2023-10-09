// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GLotteryOrderDao is the data access object for table g_lottery_order.
type GLotteryOrderDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns GLotteryOrderColumns // columns contains all the column names of Table for convenient usage.
}

// GLotteryOrderColumns defines and stores column names for table g_lottery_order.
type GLotteryOrderColumns struct {
	OrderNo    string // 编号
	Uid        string // 用户编号
	Account    string // 用户账号
	CreatedAt  string // 时间
	GameCode   string // 游戏编号
	GameName   string // 游戏名称
	Amount     string // 投注金额
	Profit     string // 赢利
	CalcAt     string // 结算时间
	Status     string // 状态 0 撤单，1:投注成功，2：中奖，3：未中奖
	Odds       string // 赔率
	Fee        string // 手续费
	OpenResult string // 开奖结果
	PlayCode   string // PlayId
	GameIssue  string // 期号
	Pid        string //
	ParentPath string //
}

// gLotteryOrderColumns holds the columns for table g_lottery_order.
var gLotteryOrderColumns = GLotteryOrderColumns{
	OrderNo:    "order_no",
	Uid:        "uid",
	Account:    "account",
	CreatedAt:  "created_at",
	GameCode:   "game_code",
	GameName:   "game_name",
	Amount:     "amount",
	Profit:     "profit",
	CalcAt:     "calc_at",
	Status:     "status",
	Odds:       "odds",
	Fee:        "fee",
	OpenResult: "open_result",
	PlayCode:   "play_code",
	GameIssue:  "game_issue",
	Pid:        "pid",
	ParentPath: "parent_path",
}

// NewGLotteryOrderDao creates and returns a new DAO object for table data access.
func NewGLotteryOrderDao() *GLotteryOrderDao {
	return &GLotteryOrderDao{
		group:   "default",
		table:   "g_lottery_order",
		columns: gLotteryOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GLotteryOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GLotteryOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GLotteryOrderDao) Columns() GLotteryOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GLotteryOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GLotteryOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GLotteryOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
