package model

type UserInfo struct {
	Uid int64 `json:"uid"`

	Account    string `json:"account"`
	ClientIP   string `json:"clientIp"`
	Lang       string `json:"lang"`
	Pid        int64  `json:"pid"`
	ParentPath string `json:"parentPath"`
	Phone      string `json:"phone"`
}
