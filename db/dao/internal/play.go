// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayDao is the data access object for table g_play.
type PlayDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns PlayColumns // columns contains all the column names of Table for convenient usage.
}

// PlayColumns defines and stores column names for table g_play.
type PlayColumns struct {
	Code          string // 玩法编号
	Name          string // 玩法名
	Status        string // 状态
	Odds          string // 赔率
	Probabilities string // 概率
	PlayType      string // 玩法类型
	BetMin        string // 最小投注金额
	BetMax        string // 最大投注金额
	Number        string //
}

// playColumns holds the columns for table g_play.
var playColumns = PlayColumns{
	Code:          "code",
	Name:          "name",
	Status:        "status",
	Odds:          "odds",
	Probabilities: "probabilities",
	PlayType:      "play_type",
	BetMin:        "bet_min",
	BetMax:        "bet_max",
	Number:        "number",
}

// NewPlayDao creates and returns a new DAO object for table data access.
func NewPlayDao() *PlayDao {
	return &PlayDao{
		group:   "default",
		table:   "g_play",
		columns: playColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PlayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PlayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PlayDao) Columns() PlayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PlayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PlayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PlayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
