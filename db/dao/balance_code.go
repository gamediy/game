// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalBalanceCodeDao is internal type for wrapping internal DAO implements.
type internalBalanceCodeDao = *internal.BalanceCodeDao

// balanceCodeDao is the data access object for table u_balance_code.
// You can define custom methods on it to extend its functionality as you wish.
type balanceCodeDao struct {
	internalBalanceCodeDao
}

var (
	// BalanceCode is globally public accessible object for table u_balance_code operations.
	BalanceCode = balanceCodeDao{
		internal.NewBalanceCodeDao(),
	}
)

// Fill with you ideas below.