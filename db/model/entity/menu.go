// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id        int64       `json:"id"        description:""`
	Pid       int64       `json:"pid"       description:""`
	Name      string      `json:"name"      description:""`
	NameZh    string      `json:"nameZh"    description:""`
	NameJa    string      `json:"nameJa"    description:""`
	Url       string      `json:"url"       description:""`
	Desc      string      `json:"desc"      description:""`
	Sort      float64     `json:"sort"      description:""`
	Type      int64       `json:"type"      description:"1 group 2 menu"`
	Status    int64       `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
