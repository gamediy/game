// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Dict is the golang structure for table dict.
type Dict struct {
	Id        uint64      `json:"id"        description:""`
	Title     string      `json:"title"     description:""`
	K         string      `json:"k"         description:""`
	V         string      `json:"v"         description:""`
	Desc      string      `json:"desc"      description:""`
	Group     string      `json:"group"     description:""`
	Status    int         `json:"status"    description:""`
	Type      int         `json:"type"      description:"1 文本，2 img"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
