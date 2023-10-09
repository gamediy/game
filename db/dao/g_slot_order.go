// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalGSlotOrderDao is internal type for wrapping internal DAO implements.
type internalGSlotOrderDao = *internal.GSlotOrderDao

// gSlotOrderDao is the data access object for table g_slot_order.
// You can define custom methods on it to extend its functionality as you wish.
type gSlotOrderDao struct {
	internalGSlotOrderDao
}

var (
	// GSlotOrder is globally public accessible object for table g_slot_order operations.
	GSlotOrder = gSlotOrderDao{
		internal.NewGSlotOrderDao(),
	}
)

// Fill with you ideas below.