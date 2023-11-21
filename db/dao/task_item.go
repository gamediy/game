// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalTaskItemDao is internal type for wrapping internal DAO implements.
type internalTaskItemDao = *internal.TaskItemDao

// taskItemDao is the data access object for table a_task_item.
// You can define custom methods on it to extend its functionality as you wish.
type taskItemDao struct {
	internalTaskItemDao
}

var (
	// TaskItem is globally public accessible object for table a_task_item operations.
	TaskItem = taskItemDao{
		internal.NewTaskItemDao(),
	}
)

// Fill with you ideas below.
