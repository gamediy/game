// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameDao is the data access object for table g_game.
type GameDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns GameColumns // columns contains all the column names of Table for convenient usage.
}

// GameColumns defines and stores column names for table g_game.
type GameColumns struct {
	Code            string // 游戏编号
	Name            string // 名称
	Status          string // 状态0：关闭，1开始
	CreatedAt       string // 创建时间
	StartTime       string // 开始时间
	EndTime         string // 结束时间
	TotalIssue      string // 一共有多少期
	IntervalSeconds string // 每期间隔秒数
	GameType        string // code
	Sort            string // 排序
	CategoryCode    string // 分类
	Number          string //
	Length          string // 号码长度
	RandomRange     string // 随机范围
	CloseSeconds    string // 封盘秒数
	PlayType        string // 玩法类型
	ThirdId         string //
	ThirdGameCode   string //
}

// gameColumns holds the columns for table g_game.
var gameColumns = GameColumns{
	Code:            "code",
	Name:            "name",
	Status:          "status",
	CreatedAt:       "created_at",
	StartTime:       "start_time",
	EndTime:         "end_time",
	TotalIssue:      "total_issue",
	IntervalSeconds: "interval_seconds",
	GameType:        "game_type",
	Sort:            "sort",
	CategoryCode:    "category_code",
	Number:          "number",
	Length:          "length",
	RandomRange:     "random_range",
	CloseSeconds:    "close_seconds",
	PlayType:        "play_type",
	ThirdId:         "third_id",
	ThirdGameCode:   "third_game_code",
}

// NewGameDao creates and returns a new DAO object for table data access.
func NewGameDao() *GameDao {
	return &GameDao{
		group:   "default",
		table:   "g_game",
		columns: gameColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameDao) Columns() GameColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
