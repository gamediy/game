package mailbox_svc

import (
	"context"
	"game/db/dao"
	"game/db/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

type SentMailBox struct {
	Title    string
	Content  string
	Receiver int64
}

func (m *SentMailBox) Exec(ctx context.Context) error {
	d := entity.Mailbox{}
	d.Title = m.Title
	d.Content = m.Content
	d.Receiver = m.Receiver
	d.Read = 0
	time := gtime.Now()
	d.ReadTime = time
	d.ReadStart = time
	if _, err := dao.Mailbox.Ctx(ctx).Insert(d); err != nil {
		return err
	}
	return nil
}
