package user

import (
	"context"
	"game/app/user/api/user/sys"
	syssvc "game/app/user/internal/service/sys"
)

func (*Controller) GetAnnouncement(ctx context.Context, _ *sys.AnnouncementReq) (res *sys.AnnouncementRes, err error) {
	return syssvc.GetAnnouncement(ctx)
}
