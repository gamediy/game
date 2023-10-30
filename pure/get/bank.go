package get

import (
	"context"
	"fmt"
	"game/db/dao"
	"game/db/model/entity"
)

func Bank(ctx context.Context, id interface{}) (*entity.Bank, error) {
	var d entity.Bank
	_ = dao.Bank.Ctx(ctx).Scan(&d, "id", id)
	if d.Id == 0 {
		return nil, fmt.Errorf("bank not found")
	}
	return &d, nil
}
