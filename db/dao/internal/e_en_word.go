// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EEnWordDao is the data access object for table e_en_word.
type EEnWordDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns EEnWordColumns // columns contains all the column names of Table for convenient usage.
}

// EEnWordColumns defines and stores column names for table e_en_word.
type EEnWordColumns struct {
	Id           string //
	Pid          string //
	HasKids      string //
	PartOfSpeech string //
	NounType     string //
	Gender       string //
	Parents      string //
	Subject      string //
	Word         string //
	Zh           string //
	Level        string //
	Deep         string //
	Status       string //
	En           string //
	Root         string //
	Prefix       string //
	Suffix       string //
	Desc         string //
	DemoSentence string //
	Img          string //
	CreatedAt    string //
	UpdatedAt    string //
}

// eEnWordColumns holds the columns for table e_en_word.
var eEnWordColumns = EEnWordColumns{
	Id:           "id",
	Pid:          "pid",
	HasKids:      "has_kids",
	PartOfSpeech: "part_of_speech",
	NounType:     "noun_type",
	Gender:       "gender",
	Parents:      "parents",
	Subject:      "subject",
	Word:         "word",
	Zh:           "zh",
	Level:        "level",
	Deep:         "deep",
	Status:       "status",
	En:           "en",
	Root:         "root",
	Prefix:       "prefix",
	Suffix:       "suffix",
	Desc:         "desc",
	DemoSentence: "demo_sentence",
	Img:          "img",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewEEnWordDao creates and returns a new DAO object for table data access.
func NewEEnWordDao() *EEnWordDao {
	return &EEnWordDao{
		group:   "default",
		table:   "e_en_word",
		columns: eEnWordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EEnWordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EEnWordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EEnWordDao) Columns() EEnWordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EEnWordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EEnWordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EEnWordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
