// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MailboxAlreadyRead is the golang structure of table u_mailbox_already_read for DAO operations like Where/Data.
type MailboxAlreadyRead struct {
	g.Meta    `orm:"table:u_mailbox_already_read, do:true"`
	Id        interface{} //
	MailBoxId interface{} //
	Uid       interface{} //
	CreatedAt *gtime.Time //
}
