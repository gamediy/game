package logic

import (
	"context"
	"game/app/slot/internal/logic/pt10"
	"game/db/model/entity"
)

var (
	GameSpin map[int32]func(ctx context.Context, play entity.Play, game entity.Game) (res []int32, bet string, err error)
)

func GameSpinInit() {
	GameSpin = make(map[int32]func(ctx context.Context, play entity.Play, game entity.Game) (res []int32, bet string, err error))
	GameSpin[10] = pt10.Spin
}
