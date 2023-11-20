// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TBetDao is the data access object for table t_bet.
type TBetDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TBetColumns // columns contains all the column names of Table for convenient usage.
}

// TBetColumns defines and stores column names for table t_bet.
type TBetColumns struct {
	Id           string //
	Uid          string //
	Account      string //
	GameTitle    string // 自己的
	GameCode     string // 自己的
	ThirdId      string // 三方ID
	ThirdName    string // 三方名称
	ThirdData    string // 三方数据
	Amount       string // 下注额
	ThirdOrderNo string // 三方订单号
	Status       string // 0:下注成功，1：中奖，2：未中奖，3：未结算
	CalcTime     string // 结算时间
	CreatedAt    string //
}

// tBetColumns holds the columns for table t_bet.
var tBetColumns = TBetColumns{
	Id:           "id",
	Uid:          "uid",
	Account:      "account",
	GameTitle:    "game_title",
	GameCode:     "game_code",
	ThirdId:      "third_id",
	ThirdName:    "third_name",
	ThirdData:    "third_data",
	Amount:       "amount",
	ThirdOrderNo: "third_order_no",
	Status:       "status",
	CalcTime:     "calc_time",
	CreatedAt:    "created_at",
}

// NewTBetDao creates and returns a new DAO object for table data access.
func NewTBetDao() *TBetDao {
	return &TBetDao{
		group:   "default",
		table:   "t_bet",
		columns: tBetColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TBetDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TBetDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TBetDao) Columns() TBetColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TBetDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TBetDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TBetDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
