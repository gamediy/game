package ws

import (
	"game/model"
)

type Controller interface {
	Controller(client *Client, msg *model.Message) (*model.Message, error)
}
