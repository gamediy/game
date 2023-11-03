package deposit_svc

import (
	"context"
	"game/app/user/api/user/deposit"
	"game/db/dao"
	"game/pure/get"
)

func ListDeposit(ctx context.Context, req *deposit.ListDepositReq) (*deposit.ListDepositRes, error) {
	var res = &deposit.ListDepositRes{}
	res.List = make([]*deposit.DepositItem, 0)
	db := dao.Deposit.Ctx(ctx).Where("uid", req.Uid)
	count, err := db.Count()
	if err != nil {
		return nil, err
	}
	res.Total = int64(count)
	page := req.Page
	size := req.Size
	if size <= 0 || size >= 100 {
		size = 10
	}
	if err = db.Page(int(page), int(size)).
		Fields("uid,id order_no,amount,status,currency,protocol,amount_item_id,transfer_order_no,transfer_img,created_at,status,status_remark").
		Scan(&res.List); err != nil {
		return nil, err
	}
	if len(res.List) != 0 {
		for _, i := range res.List {
			amountItem, err := get.AmountItemFromCache(ctx, i.AmountItemId)
			if err != nil {
				continue
			}
			i.Logo = get.ImgPrefix() + amountItem.Logo
		}
	}
	return res, nil
}
