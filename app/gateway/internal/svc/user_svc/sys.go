package user_svc

import (
	"context"
	"game/app/user/api/user/sys"
)

func (s userSvc) GetAnnouncement(ctx context.Context, s2 *sys.AnnouncementReq) (interface{}, error) {
	return sysClient.GetAnnouncement(ctx, s2)
}
