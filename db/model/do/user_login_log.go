// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLoginLog is the golang structure of table u_user_login_log for DAO operations like Where/Data.
type UserLoginLog struct {
	g.Meta    `orm:"table:u_user_login_log, do:true"`
	Id        interface{} //
	Uid       interface{} //
	Account   interface{} //
	Ip        interface{} //
	CreatedAt *gtime.Time //
}
