// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsVerification is the golang structure of table u_sms_verification for DAO operations like Where/Data.
type SmsVerification struct {
	g.Meta     `orm:"table:u_sms_verification, do:true"`
	Id         interface{} //
	Phone      interface{} //
	Code       interface{} //
	Times      interface{} // number of verifications
	Expiration *gtime.Time //
	CreatedAt  *gtime.Time //
}
