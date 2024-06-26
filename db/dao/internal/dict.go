// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DictDao is the data access object for table s_dict.
type DictDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns DictColumns // columns contains all the column names of Table for convenient usage.
}

// DictColumns defines and stores column names for table s_dict.
type DictColumns struct {
	Id        string //
	Title     string //
	K         string //
	V         string //
	Desc      string //
	Group     string //
	Status    string //
	Type      string // 1 文本，2 img
	CreatedAt string //
	UpdatedAt string //
}

// dictColumns holds the columns for table s_dict.
var dictColumns = DictColumns{
	Id:        "id",
	Title:     "title",
	K:         "k",
	V:         "v",
	Desc:      "desc",
	Group:     "group",
	Status:    "status",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewDictDao creates and returns a new DAO object for table data access.
func NewDictDao() *DictDao {
	return &DictDao{
		group:   "default",
		table:   "s_dict",
		columns: dictColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DictDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DictDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DictDao) Columns() DictColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DictDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DictDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DictDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
