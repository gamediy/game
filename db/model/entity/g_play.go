// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GPlay is the golang structure for table g_play.
type GPlay struct {
	Code          int     `json:"code"          description:"玩法编号"`
	Name          string  `json:"name"          description:"玩法名"`
	Status        int     `json:"status"        description:"状态"`
	Type          string  `json:"type"          description:"类型"`
	Odds          float64 `json:"odds"          description:"赔率"`
	Probabilities float64 `json:"probabilities" description:"概率"`
	GameCode      int     `json:"gameCode"      description:"游戏编号"`
	BetInx        float64 `json:"betInx"        description:"最小投注金额"`
	BetMax        float64 `json:"betMax"        description:"最大投注金额"`
}
