// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        int64       `json:"id"        description:""`
	Rid       int64       `json:"rid"       description:""`
	Account   string      `json:"account"   description:""`
	Pass      string      `json:"pass"      description:""`
	Nickname  string      `json:"nickname"  description:""`
	Desc      string      `json:"desc"      description:""`
	Email     string      `json:"email"     description:""`
	Phone     string      `json:"phone"     description:""`
	Status    int64       `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	KeySecret string      `json:"keySecret" description:"Google-Authenticator key secret"`
}
