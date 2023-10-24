// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        int64       `json:"id"        description:""`
	Name      string      `json:"name"      description:""`
	NameZh    string      `json:"nameZh"    description:""`
	NameJa    string      `json:"nameJa"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	Class     string      `json:"class"     description:""`
}
