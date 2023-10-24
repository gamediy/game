// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WalletReportDay is the golang structure for table wallet_report_day.
type WalletReportDay struct {
	Id        int64       `json:"id"        description:""`
	Uid       int64       `json:"uid"       description:""`
	Account   string      `json:"account"   description:""`
	TransCode int         `json:"transCode" description:""`
	Date      *gtime.Time `json:"date"      description:""`
	Amount    float64     `json:"amount"    description:""`
	Count     int         `json:"count"     description:""`
	Title     string      `json:"title"     description:""`
}
