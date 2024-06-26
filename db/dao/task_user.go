// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalTaskUserDao is internal type for wrapping internal DAO implements.
type internalTaskUserDao = *internal.TaskUserDao

// taskUserDao is the data access object for table a_task_user.
// You can define custom methods on it to extend its functionality as you wish.
type taskUserDao struct {
	internalTaskUserDao
}

var (
	// TaskUser is globally public accessible object for table a_task_user operations.
	TaskUser = taskUserDao{
		internal.NewTaskUserDao(),
	}
)

// Fill with you ideas below.
