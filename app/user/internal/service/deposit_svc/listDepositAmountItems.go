package deposit_svc

import (
	"context"
	"fmt"
	"game/app/user/api/user/deposit"
	"game/db/model/entity"
	"game/pure/get"
	"game/utility/blockchain/tron/client"

	"game/db/dao"
)

func ListDepositAmountItems(ctx context.Context, uid int64) (*deposit.DepositAmountItemsRes, error) {
	var (
		res      deposit.DepositAmountItemsRes
		list     = make([]*deposit.AmountItem, 0)
		category = make([]entity.AmountCategory, 0)
	)
	res.List = make([]*deposit.DepositBox, 0)
	res.Tips = "payment prompt get configuration file"
	userInfo, err := get.UserById(ctx, uid)
	if err != nil {
		return nil, err
	}
	dao.AmountItem.Ctx(ctx).Scan(&list, "status=? and type=?", 1, "Deposit")
	dao.AmountCategory.Ctx(ctx).Scan(&category, "status", 1)
	for _, category := range category {
		depositList := &deposit.DepositBox{}
		depositList.Title = category.Title
		depositList.CategoryId = int64(category.Id)

		for _, item := range list {
			if item.AmountCategoryId == int64(category.Id) {
				if item.Protocol == "Trc20" && item.Address == "" {
					account := entity.DigitalAccount{}
					dao.DigitalAccount.Ctx(ctx).Where("status=1 and uid=?", uid).Scan(&account)
					if account.Address == "" {
						key, addr := client.TronGrpcClient.GenerateKey()
						account.Address = addr
						account.PrivateKey = key
						account.Account = userInfo.Account
						account.Uid = int(userInfo.Id)
						account.Status = 1
						account.Net = "TRON"
						dao.DigitalAccount.Ctx(ctx).Data(&account).Insert()
					}
					item.Address = account.Address
				}

				depositList.Item = append(depositList.Item, &deposit.AmountItem{
					Id:           item.Id,
					Title:        item.Title,
					Protocol:     item.Protocol,
					Logo:         fmt.Sprint(get.ImgPrefix(), item.Logo),
					Currency:     item.Currency,
					Address:      item.Address,
					Max:          item.Max,
					Min:          item.Min,
					Detail:       item.Detail,
					BankId:       item.BankId,
					ExchangeRate: item.ExchangeRate,
					SelectMoney:  item.SelectMoney,
				})
			}

		}
		if len(depositList.Item) > 0 {
			res.List = append(res.List, depositList)
		}

	}

	return &res, nil
}
