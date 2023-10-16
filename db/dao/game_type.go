// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalGameTypeDao is internal type for wrapping internal DAO implements.
type internalGameTypeDao = *internal.GameTypeDao

// gameTypeDao is the data access object for table g_game_type.
// You can define custom methods on it to extend its functionality as you wish.
type gameTypeDao struct {
	internalGameTypeDao
}

var (
	// GameType is globally public accessible object for table g_game_type operations.
	GameType = gameTypeDao{
		internal.NewGameTypeDao(),
	}
)

// Fill with you ideas below.
