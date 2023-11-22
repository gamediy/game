// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TransactionType is the golang structure for table transaction_type.
type TransactionType struct {
	Id        int64       `json:"id"        description:""`
	Title     string      `json:"title"     description:""`
	Code      int         `json:"code"      description:""`
	Remark    string      `json:"remark"    description:""`
	Type      int         `json:"type"      description:""`
	Class     string      `json:"class"     description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
