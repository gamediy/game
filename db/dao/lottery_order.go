// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalLotteryOrderDao is internal type for wrapping internal DAO implements.
type internalLotteryOrderDao = *internal.LotteryOrderDao

// lotteryOrderDao is the data access object for table g_lottery_order.
// You can define custom methods on it to extend its functionality as you wish.
type lotteryOrderDao struct {
	internalLotteryOrderDao
}

var (
	// LotteryOrder is globally public accessible object for table g_lottery_order operations.
	LotteryOrder = lotteryOrderDao{
		internal.NewLotteryOrderDao(),
	}
)

// Fill with you ideas below.
