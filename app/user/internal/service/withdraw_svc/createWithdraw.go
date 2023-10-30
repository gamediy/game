package withdraw_svc

type CreateWithdraw struct {
	AmountItemId      int
	Amount            float64
	WithdrawAccountId int
}
