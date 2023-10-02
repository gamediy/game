package user_svc

import (
	"context"
	"fmt"
	"game/app/user/api/withdraw/withdraw"
)

func (userSvc) SubmitWithdraw(ctx context.Context) {

	submit, err := withdrawClient.Submit(ctx, &withdraw.SubmitRequest{Amount: 1})
	if err == nil {
		fmt.Println(submit)
	}
}
