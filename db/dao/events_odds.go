// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalEventsOddsDao is internal type for wrapping internal DAO implements.
type internalEventsOddsDao = *internal.EventsOddsDao

// eventsOddsDao is the data access object for table p_events_odds.
// You can define custom methods on it to extend its functionality as you wish.
type eventsOddsDao struct {
	internalEventsOddsDao
}

var (
	// EventsOdds is globally public accessible object for table p_events_odds operations.
	EventsOdds = eventsOddsDao{
		internal.NewEventsOddsDao(),
	}
)

// Fill with you ideas below.