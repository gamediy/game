package pt10

import (
	"context"
	"encoding/json"
	"fmt"
	"game/db/model/entity"
	"math/rand"
)

var (
	gameNumber = map[int32]map[int32]string{
		1000: {
			1: "7",
			2: "Apple",
			3: "Banana",
			4: "Bar",
			5: "Bell",
			6: "Grape",
			7: "Cherry",
			8: "Orange",
		},
	}
)

func Spin(ctx context.Context, play entity.Play, game entity.Game) (res []int32, bet string, err error) {

	res = []int32{}
	if play.Number == "1000" { //生成水果连线
		numbers := []int32{2, 3, 5, 6, 7, 8}
		randomNumber := numbers[rand.Intn(len(numbers))]
		res = []int32{randomNumber, randomNumber, randomNumber}

	} else if play.Number == "1001" {
		numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8}
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		res = numbers[:3]
	}
	json.Unmarshal([]byte(play.Number), &res)
	m, ok := gameNumber[int32(game.Code)]
	if ok {
		bet = fmt.Sprintf("%s,%s,%s", m[res[0]], m[res[1]], m[res[2]])
	}

	return res, bet, nil
}
