// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalApiDao is internal type for wrapping internal DAO implements.
type internalApiDao = *internal.ApiDao

// apiDao is the data access object for table s_api.
// You can define custom methods on it to extend its functionality as you wish.
type apiDao struct {
	internalApiDao
}

var (
	// Api is globally public accessible object for table s_api operations.
	Api = apiDao{
		internal.NewApiDao(),
	}
)

// Fill with you ideas below.
