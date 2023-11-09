package gamesvc

import (
	"context"
	"game/app/game/api/game/game"
	"game/db/dao"
	"game/pure/get"
)

func ListBanner(ctx context.Context) (*game.BannerRes, error) {
	var res = &game.BannerRes{}
	res.List = make([]*game.BannerItem, 0)
	err := dao.Banner.Ctx(ctx).Scan(&res.List)
	for _, i := range res.List {
		i.Image = get.ImgPrefix() + i.Image
	}
	return res, err
}
