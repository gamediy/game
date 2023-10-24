// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WalletReportDay is the golang structure of table w_wallet_report_day for DAO operations like Where/Data.
type WalletReportDay struct {
	g.Meta    `orm:"table:w_wallet_report_day, do:true"`
	Id        interface{} //
	Uid       interface{} //
	Account   interface{} //
	TransCode interface{} //
	Date      *gtime.Time //
	Amount    interface{} //
	Count     interface{} //
	Title     interface{} //
}
