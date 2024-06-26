// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Api is the golang structure of table s_api for DAO operations like Where/Data.
type Api struct {
	g.Meta    `orm:"table:s_api, do:true"`
	Id        interface{} //
	Group     interface{} //
	Url       interface{} //
	Method    interface{} // 1 get 2 post 3 put 4 delete
	Desc      interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
