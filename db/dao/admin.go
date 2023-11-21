// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalAdminDao is internal implement for wrapping internal DAO implements.
type internalAdminDao = *internal.AdminDao

// adminDao is the data access object for table s_admin.
// You can define custom methods on it to extend its functionality as you wish.
type adminDao struct {
	internalAdminDao
}

var (
	// Admin is globally public accessible object for table s_admin operations.
	Admin = adminDao{
		internal.NewAdminDao(),
	}
)

// Fill with you ideas below.
