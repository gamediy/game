// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure for table words.
type Words struct {
	Id        int64       `json:"id"        description:""`
	K         string      `json:"k"         description:""`
	En        string      `json:"en"        description:""`
	Zh        string      `json:"zh"        description:""`
	Ja        string      `json:"ja"        description:""`
	Group     string      `json:"group"     description:""`
	Desc      string      `json:"desc"      description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}
