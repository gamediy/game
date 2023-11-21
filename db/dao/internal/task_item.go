// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskItemDao is the data access object for table a_task_item.
type TaskItemDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TaskItemColumns // columns contains all the column names of Table for convenient usage.
}

// TaskItemColumns defines and stores column names for table a_task_item.
type TaskItemColumns struct {
	Id             string // 编号
	Title          string // 标题
	LevelNum       string // 级别
	Status         string // 状态 0
	Type           string // 类型 1签到
	TypeCode       string // 类型值
	EndValue       string // 结束值
	TimeType       string // 时间类型
	StartTime      string // 开始时间
	EndTime        string // 结束时间
	RewardRule     string // 奖利规则
	GiveType       string // 领取方式   1,所有用户,2 VIP,3,userId,4,level
	GiveValue      string // 领取值
	CurrentValue   string // 当前值
	StartValue     string // 开始值
	TaskCategoryId string // 分类编号
	Remark         string // 备注
	TaskId         string // 任务编号
}

// taskItemColumns holds the columns for table a_task_item.
var taskItemColumns = TaskItemColumns{
	Id:             "id",
	Title:          "title",
	LevelNum:       "level_num",
	Status:         "status",
	Type:           "type",
	TypeCode:       "type_code",
	EndValue:       "end_value",
	TimeType:       "time_type",
	StartTime:      "start_time",
	EndTime:        "end_time",
	RewardRule:     "reward_rule",
	GiveType:       "give_type",
	GiveValue:      "give_value",
	CurrentValue:   "current_value",
	StartValue:     "start_value",
	TaskCategoryId: "task_category_id",
	Remark:         "remark",
	TaskId:         "task_id",
}

// NewTaskItemDao creates and returns a new DAO object for table data access.
func NewTaskItemDao() *TaskItemDao {
	return &TaskItemDao{
		group:   "default",
		table:   "a_task_item",
		columns: taskItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskItemDao) Columns() TaskItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
