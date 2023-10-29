// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Payment is the golang structure for table payment.
type Payment struct {
	Merchant string `json:"merchant" description:"商户"`
	Key      string `json:"key"      description:""`
	Name     string `json:"name"     description:""`
	Status   int    `json:"status"   description:""`
	Url      string `json:"url"      description:""`
	Code     int    `json:"code"     description:"编号"`
}
