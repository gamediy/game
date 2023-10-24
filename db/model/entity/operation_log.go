// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OperationLog is the golang structure for table operation_log.
type OperationLog struct {
	Id        int64       `json:"id"        description:""`
	Uid       int64       `json:"uid"       description:""`
	Account   string      `json:"account"   description:""`
	RoleName  string      `json:"roleName"  description:""`
	Req       string      `json:"req"       description:""`
	Res       string      `json:"res"       description:""`
	Method    string      `json:"method"    description:""`
	Uri       string      `json:"uri"       description:""`
	Ip        string      `json:"ip"        description:""`
	UseTime   int64       `json:"useTime"   description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}
