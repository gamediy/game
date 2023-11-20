// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ATaskCategory is the golang structure for table a_task_category.
type ATaskCategory struct {
	Id      int64  `json:"id"      description:"编号"`
	Title   string `json:"title"   description:"标题"`
	Content string `json:"content" description:"内容"`
	Status  int    `json:"status"  description:"状态"`
	Icon    string `json:"icon"    description:""`
	Images  string `json:"images"  description:""`
}
