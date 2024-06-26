// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Banner is the golang structure for table banner.
type Banner struct {
	Id        uint64      `json:"id"        description:""`
	Title     string      `json:"title"     description:""`
	Image     string      `json:"image"     description:""`
	Link      string      `json:"link"      description:""`
	Desc      string      `json:"desc"      description:""`
	Sort      int         `json:"sort"      description:""`
	Status    uint        `json:"status"    description:"1open 2close"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}
