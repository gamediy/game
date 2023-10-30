package withdraw_svc

import (
	"context"
	"fmt"
	"game/db/dao"
	"game/pure/get"
	"game/utility/utils/xpwd"
	"github.com/gogf/gf/v2/frame/g"
)

type SetPayPass struct {
	Uid  int64
	Pass string
}

func (m *SetPayPass) Exec(ctx context.Context) error {
	u, err := get.UserById(ctx, m.Uid)
	if err != nil {
		return err
	}
	if u.PayPass != "" {
		return fmt.Errorf("pay pass has been set")
	}
	d := g.Map{dao.User.Columns().PayPass: xpwd.GenPwd(m.Pass)}
	if _, err = dao.User.Ctx(ctx).Update(d, "id", u.Id); err != nil {
		return err
	}
	return nil
}
