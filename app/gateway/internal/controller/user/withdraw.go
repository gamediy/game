package user

import (
	"fmt"
	withdraw_pb "game/app/user/pb/withdraw"
)

type withdraw struct {
}

func (withdraw) Submit() {

	var (
		client = withdraw_pb.NewWithdrawServiceClient(conn)
	)
	submit, err := client.Submit(ctx, &withdraw_pb.SubmitRequest{Amount: 1})
	if err == nil {
		fmt.Println(submit)
	}
}
