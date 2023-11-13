package v1

import (
	"game/utility/utils/xtime"
	"github.com/gogf/gf/v2/frame/g"
)

func SetStatusOk() *Status {
	return &Status{Code: "0", Message: "Success", Datetime: xtime.NowForISO8601()}
}

type Status struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Datetime string `json:"datetime"`
}
type CheckPlayerReq struct {
	g.Meta  `tags:"check player" method:"get" path:"/player/check/:account"`
	Account string `json:"account"`
}
type CheckPlayerRes struct {
	Data   bool    `json:"data"`
	Status *Status `json:"status"`
}

type BalanceReq struct {
	g.Meta  `method:"get" path:"/transaction/balance/:account"`
	Account string `json:"account"`
}
type BalanceRes struct {
	Data struct {
		Balance  float64 `json:"balance"`
		Currency string  `json:"currency"`
	}
	Status *Status `json:"status"`
}
