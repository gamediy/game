// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Mailbox is the golang structure of table u_mailbox for DAO operations like Where/Data.
type Mailbox struct {
	g.Meta    `orm:"table:u_mailbox, do:true"`
	Id        interface{} //
	Title     interface{} //
	CreatedAt *gtime.Time //
	Content   interface{} //
	Receiver  interface{} // 接收者（所有（0），或者指定用户ID）
	Read      interface{} // 0：未读，1：已读
	ReadTime  *gtime.Time // 已读时间
	ReadStart *gtime.Time // 开始阅读时间
	Type      interface{} // 0 普通 1 赠送
}
