// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskDao is the data access object for table a_task.
type TaskDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TaskColumns // columns contains all the column names of Table for convenient usage.
}

// TaskColumns defines and stores column names for table a_task.
type TaskColumns struct {
	Id             string // 编号
	TaskCategoryId string // 分类编号
	Title          string // 标题
	LevelNum       string // 关卡数
	Status         string // 状态0：未开始，1:进行中，2：结束
	Type           string // 类型   1:签到
	StartTime      string // 开始时间
	EndTime        string // 结束时间
	TimeType       string // 1:每天，2：每周，3每月,4:每小时 99:永久
	GiveType       string // 送取类型    1,所有用户,2 VIP,3,userId,4,level
	GiveValue      string // 数据
	TimeValue      string //
	Remark         string //
}

// taskColumns holds the columns for table a_task.
var taskColumns = TaskColumns{
	Id:             "id",
	TaskCategoryId: "task_category_id",
	Title:          "title",
	LevelNum:       "level_num",
	Status:         "status",
	Type:           "implement",
	StartTime:      "start_time",
	EndTime:        "end_time",
	TimeType:       "time_type",
	GiveType:       "give_type",
	GiveValue:      "give_value",
	TimeValue:      "time_value",
	Remark:         "remark",
}

// NewTaskDao creates and returns a new DAO object for table data access.
func NewTaskDao() *TaskDao {
	return &TaskDao{
		group:   "default",
		table:   "a_task",
		columns: taskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskDao) Columns() TaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
