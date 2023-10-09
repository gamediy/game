// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GPlayDao is the data access object for table g_play.
type GPlayDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns GPlayColumns // columns contains all the column names of Table for convenient usage.
}

// GPlayColumns defines and stores column names for table g_play.
type GPlayColumns struct {
	Code          string // 玩法编号
	Name          string // 玩法名
	Status        string // 状态
	Type          string // 类型
	Odds          string // 赔率
	Probabilities string // 概率
	GameCode      string // 游戏编号
	BetInx        string // 最小投注金额
	BetMax        string // 最大投注金额
}

// gPlayColumns holds the columns for table g_play.
var gPlayColumns = GPlayColumns{
	Code:          "code",
	Name:          "name",
	Status:        "status",
	Type:          "type",
	Odds:          "odds",
	Probabilities: "probabilities",
	GameCode:      "game_code",
	BetInx:        "bet_inx",
	BetMax:        "bet_max",
}

// NewGPlayDao creates and returns a new DAO object for table data access.
func NewGPlayDao() *GPlayDao {
	return &GPlayDao{
		group:   "default",
		table:   "g_play",
		columns: gPlayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GPlayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GPlayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GPlayDao) Columns() GPlayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GPlayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GPlayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GPlayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
