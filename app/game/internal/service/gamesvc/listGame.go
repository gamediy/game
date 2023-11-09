package gamesvc

import (
	"context"
	"game/app/game/api/game/game"
	"game/db/dao"
)

func ListGame(ctx context.Context, code string) (*game.ListGameRes, error) {
	res := &game.ListGameRes{}
	res.List = make([]*game.GameItem, 0)
	db := dao.Game.Ctx(ctx)
	if code != "" {
		db = db.Where("category_code", code)
	}
	err := db.Order("sort desc").Scan(&res.List)
	return res, err

}
