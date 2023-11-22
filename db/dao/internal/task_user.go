// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskUserDao is the data access object for table a_task_user.
type TaskUserDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TaskUserColumns // columns contains all the column names of Table for convenient usage.
}

// TaskUserColumns defines and stores column names for table a_task_user.
type TaskUserColumns struct {
	Id             string //
	Uid            string // 用户ID
	Account        string // 账号
	Status         string // ：未开始，1进行中，2 已完成
	CurrentValue   string // 当前值
	StartTime      string // 开始时间
	TaskId         string // 任务ID
	TaskItemId     string //
	TaskCategoryId string //
	EndTime        string // 结束时间
	GiveTime       string // 领取时间
	TypeCode       string // 类型编号
	Type           string // 类型
	StartValue     string // 开始值
	EndValue       string // 结束值
	LevelNum       string // 关卡数
	Title          string // 标题
	SuccessTime    string // 成功时间
}

// taskUserColumns holds the columns for table a_task_user.
var taskUserColumns = TaskUserColumns{
	Id:             "id",
	Uid:            "uid",
	Account:        "account",
	Status:         "status",
	CurrentValue:   "current_value",
	StartTime:      "start_time",
	TaskId:         "task_id",
	TaskItemId:     "task_item_id",
	TaskCategoryId: "task_category_id",
	EndTime:        "end_time",
	GiveTime:       "give_time",
	TypeCode:       "type_code",
	Type:           "type",
	StartValue:     "start_value",
	EndValue:       "end_value",
	LevelNum:       "level_num",
	Title:          "title",
	SuccessTime:    "success_time",
}

// NewTaskUserDao creates and returns a new DAO object for table data access.
func NewTaskUserDao() *TaskUserDao {
	return &TaskUserDao{
		group:   "default",
		table:   "a_task_user",
		columns: taskUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskUserDao) Columns() TaskUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
