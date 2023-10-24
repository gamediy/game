// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TransactionType is the golang structure of table w_transaction_type for DAO operations like Where/Data.
type TransactionType struct {
	g.Meta    `orm:"table:w_transaction_type, do:true"`
	Id        interface{} //
	Title     interface{} //
	Code      interface{} //
	Remark    interface{} //
	Type      interface{} //
	Class     interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
