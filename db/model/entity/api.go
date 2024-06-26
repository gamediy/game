// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Api is the golang structure for table api.
type Api struct {
	Id        uint64      `json:"id"        description:""`
	Group     string      `json:"group"     description:""`
	Url       string      `json:"url"       description:""`
	Method    string      `json:"method"    description:"1 get 2 post 3 put 4 delete"`
	Desc      string      `json:"desc"      description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
