// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Activity is the golang structure for table activity.
type Activity struct {
	Id      int    `json:"id"      description:""`
	Title   string `json:"title"   description:""`
	Type    int    `json:"type"    description:"类型"`
	Content string `json:"content" description:""`
	Images  string `json:"images"  description:""`
	Status  int    `json:"status"  description:""`
	Event   string `json:"event"   description:"事件"`
	Rule    string `json:"rule"    description:""`
}
