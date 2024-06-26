// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalWalletChangeLogDao is internal type for wrapping internal DAO implements.
type internalWalletChangeLogDao = *internal.WalletChangeLogDao

// walletChangeLogDao is the data access object for table w_wallet_change_log.
// You can define custom methods on it to extend its functionality as you wish.
type walletChangeLogDao struct {
	internalWalletChangeLogDao
}

var (
	// WalletChangeLog is globally public accessible object for table w_wallet_change_log operations.
	WalletChangeLog = walletChangeLogDao{
		internal.NewWalletChangeLogDao(),
	}
)

// Fill with you ideas below.
