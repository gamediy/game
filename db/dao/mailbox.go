// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game/db/dao/internal"
)

// internalMailboxDao is internal type for wrapping internal DAO implements.
type internalMailboxDao = *internal.MailboxDao

// mailboxDao is the data access object for table u_mailbox.
// You can define custom methods on it to extend its functionality as you wish.
type mailboxDao struct {
	internalMailboxDao
}

var (
	// Mailbox is globally public accessible object for table u_mailbox operations.
	Mailbox = mailboxDao{
		internal.NewMailboxDao(),
	}
)

// Fill with you ideas below.
