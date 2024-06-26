// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BannerDao is the data access object for table c_banner.
type BannerDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns BannerColumns // columns contains all the column names of Table for convenient usage.
}

// BannerColumns defines and stores column names for table c_banner.
type BannerColumns struct {
	Id        string //
	Title     string //
	Image     string //
	Link      string //
	Desc      string //
	Sort      string //
	Status    string // 1open 2close
	CreatedAt string //
}

// bannerColumns holds the columns for table c_banner.
var bannerColumns = BannerColumns{
	Id:        "id",
	Title:     "title",
	Image:     "image",
	Link:      "link",
	Desc:      "desc",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
}

// NewBannerDao creates and returns a new DAO object for table data access.
func NewBannerDao() *BannerDao {
	return &BannerDao{
		group:   "default",
		table:   "c_banner",
		columns: bannerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BannerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BannerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BannerDao) Columns() BannerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BannerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BannerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BannerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
