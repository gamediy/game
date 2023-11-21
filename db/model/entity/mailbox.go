// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Mailbox is the golang structure for table mailbox.
type Mailbox struct {
	Id        int         `json:"id"        description:""`
	Title     string      `json:"title"     description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	Content   string      `json:"content"   description:""`
	Receiver  int64       `json:"receiver"  description:"接收者（所有（0），或者指定用户ID）"`
	Read      int         `json:"read"      description:"0：未读，1：已读"`
	ReadTime  *gtime.Time `json:"readTime"  description:"已读时间"`
	ReadStart *gtime.Time `json:"readStart" description:"开始阅读时间"`
	Type      int         `json:"implement"      description:"0 普通 1 赠送"`
}
