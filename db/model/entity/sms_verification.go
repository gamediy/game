// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsVerification is the golang structure for table sms_verification.
type SmsVerification struct {
	Id         int64       `json:"id"         description:""`
	Phone      string      `json:"phone"      description:""`
	Code       string      `json:"code"       description:""`
	Times      int         `json:"times"      description:"number of verifications"`
	Expiration *gtime.Time `json:"expiration" description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  description:""`
}
