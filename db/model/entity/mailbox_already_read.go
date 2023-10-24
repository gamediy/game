// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MailboxAlreadyRead is the golang structure for table mailbox_already_read.
type MailboxAlreadyRead struct {
	Id        uint64      `json:"id"        description:""`
	MailBoxId int64       `json:"mailBoxId" description:""`
	Uid       int64       `json:"uid"       description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}
