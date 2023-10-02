package sync_controller

import "game/model"

type Controller interface {
	Controller(data []byte) (*model.WsMessage, error)
}
