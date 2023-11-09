package gamesvc

import (
	"context"
	"game/app/game/api/game/game"
	"game/db/dao"
	"game/pure/get"
)

func ListGameCategory(ctx context.Context) (*game.ListGameCategoryRes, error) {
	res := &game.ListGameCategoryRes{}
	res.List = make([]*game.GameCategoryItem, 0)
	err := dao.GameCategory.Ctx(ctx).Scan(&res.List)
	if err != nil {
		return res, err
	}
	for _, i := range res.List {
		i.Logo = get.ImgPrefix() + i.Logo
	}
	return res, nil
}
