package play

type Play interface {
	Won(openResult interface{}, betContent interface{}, won *Won) //结算
	Check(betContent interface{}) error                           //检测
}

type WonItem struct {
	PlayCode string
}
type Won struct {
	GameCode int32
	GameName string
	List     []WonItem
}
