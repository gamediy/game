// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Play is the golang structure for table play.
type Play struct {
	Code          int     `json:"code"          description:"玩法编号"`
	Name          string  `json:"name"          description:"玩法名"`
	Status        int     `json:"status"        description:"状态"`
	Odds          float64 `json:"odds"          description:"赔率"`
	Probabilities float64 `json:"probabilities" description:"概率"`
	PlayType      int     `json:"playType"      description:"玩法类型"`
	BetMin        float64 `json:"betMin"        description:"最小投注金额"`
	BetMax        float64 `json:"betMax"        description:"最大投注金额"`
	Number        string  `json:"number"        description:""`
}
