// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TaskCategory is the golang structure of table a_task_category for DAO operations like Where/Data.
type TaskCategory struct {
	g.Meta  `orm:"table:a_task_category, do:true"`
	Id      interface{} // 编号
	Title   interface{} // 标题
	Content interface{} // 内容
	Status  interface{} // 状态
	Icon    interface{} //
	Images  interface{} //
}
