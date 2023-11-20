// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SmsVerificationDao is the data access object for table u_sms_verification.
type SmsVerificationDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SmsVerificationColumns // columns contains all the column names of Table for convenient usage.
}

// SmsVerificationColumns defines and stores column names for table u_sms_verification.
type SmsVerificationColumns struct {
	Id         string //
	Phone      string //
	Code       string //
	Status     string // 0 wait, 1 ok 2 error
	Times      string // number of verifications
	Expiration string //
	CreatedAt  string //
}

// smsVerificationColumns holds the columns for table u_sms_verification.
var smsVerificationColumns = SmsVerificationColumns{
	Id:         "id",
	Phone:      "phone",
	Code:       "code",
	Status:     "status",
	Times:      "times",
	Expiration: "expiration",
	CreatedAt:  "created_at",
}

// NewSmsVerificationDao creates and returns a new DAO object for table data access.
func NewSmsVerificationDao() *SmsVerificationDao {
	return &SmsVerificationDao{
		group:   "default",
		table:   "u_sms_verification",
		columns: smsVerificationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SmsVerificationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SmsVerificationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SmsVerificationDao) Columns() SmsVerificationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SmsVerificationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SmsVerificationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SmsVerificationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
