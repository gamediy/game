// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskCategoryDao is the data access object for table a_task_category.
type TaskCategoryDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TaskCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// TaskCategoryColumns defines and stores column names for table a_task_category.
type TaskCategoryColumns struct {
	Id      string // 编号
	Title   string // 标题
	Content string // 内容
	Status  string // 状态
	Icon    string //
	Images  string //
}

// taskCategoryColumns holds the columns for table a_task_category.
var taskCategoryColumns = TaskCategoryColumns{
	Id:      "id",
	Title:   "title",
	Content: "content",
	Status:  "status",
	Icon:    "icon",
	Images:  "images",
}

// NewTaskCategoryDao creates and returns a new DAO object for table data access.
func NewTaskCategoryDao() *TaskCategoryDao {
	return &TaskCategoryDao{
		group:   "default",
		table:   "a_task_category",
		columns: taskCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskCategoryDao) Columns() TaskCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
