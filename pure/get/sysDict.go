package get

import (
	"backend/db/dao"
	"backend/db/model/entity"
	"context"
	"fmt"
)

func Dict(ctx context.Context, key string) (*entity.Dict, error) {
	var d entity.Dict
	_ = dao.Dict.Ctx(ctx).Scan(&d, "k", key)
	if d.Id == 0 {
		return nil, fmt.Errorf("data not found")
	}
	return &d, nil
}
