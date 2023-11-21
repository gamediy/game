// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MailboxDao is the data access object for table u_mailbox.
type MailboxDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns MailboxColumns // columns contains all the column names of Table for convenient usage.
}

// MailboxColumns defines and stores column names for table u_mailbox.
type MailboxColumns struct {
	Id        string //
	Title     string //
	CreatedAt string //
	Content   string //
	Receiver  string // 接收者（所有（0），或者指定用户ID）
	Read      string // 0：未读，1：已读
	ReadTime  string // 已读时间
	ReadStart string // 开始阅读时间
	Type      string // 0 普通 1 赠送
}

// mailboxColumns holds the columns for table u_mailbox.
var mailboxColumns = MailboxColumns{
	Id:        "id",
	Title:     "title",
	CreatedAt: "created_at",
	Content:   "content",
	Receiver:  "receiver",
	Read:      "read",
	ReadTime:  "read_time",
	ReadStart: "read_start",
	Type:      "implement",
}

// NewMailboxDao creates and returns a new DAO object for table data access.
func NewMailboxDao() *MailboxDao {
	return &MailboxDao{
		group:   "default",
		table:   "u_mailbox",
		columns: mailboxColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MailboxDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MailboxDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MailboxDao) Columns() MailboxColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MailboxDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MailboxDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MailboxDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
