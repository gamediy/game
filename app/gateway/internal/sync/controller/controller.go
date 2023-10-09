package sync_controller

import (
	"game/model"
)

var (
	ControllerC map[string]Controller
)

type Controller interface {
	Controller(data []byte) (*model.WsMessage, error)
}
