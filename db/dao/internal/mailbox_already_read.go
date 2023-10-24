// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MailboxAlreadyReadDao is the data access object for table u_mailbox_already_read.
type MailboxAlreadyReadDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns MailboxAlreadyReadColumns // columns contains all the column names of Table for convenient usage.
}

// MailboxAlreadyReadColumns defines and stores column names for table u_mailbox_already_read.
type MailboxAlreadyReadColumns struct {
	Id        string //
	MailBoxId string //
	Uid       string //
	CreatedAt string //
}

// mailboxAlreadyReadColumns holds the columns for table u_mailbox_already_read.
var mailboxAlreadyReadColumns = MailboxAlreadyReadColumns{
	Id:        "id",
	MailBoxId: "mail_box_id",
	Uid:       "uid",
	CreatedAt: "created_at",
}

// NewMailboxAlreadyReadDao creates and returns a new DAO object for table data access.
func NewMailboxAlreadyReadDao() *MailboxAlreadyReadDao {
	return &MailboxAlreadyReadDao{
		group:   "default",
		table:   "u_mailbox_already_read",
		columns: mailboxAlreadyReadColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MailboxAlreadyReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MailboxAlreadyReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MailboxAlreadyReadDao) Columns() MailboxAlreadyReadColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MailboxAlreadyReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MailboxAlreadyReadDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MailboxAlreadyReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
