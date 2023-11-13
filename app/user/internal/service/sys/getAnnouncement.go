package sys

import (
	"context"
	"game/app/user/api/user/sys"
	"game/pure/get"
)

func GetAnnouncement(ctx context.Context) (*sys.AnnouncementRes, error) {
	dict, err := get.Dict(ctx, "announcement")
	if err != nil {
		return nil, err
	}
	return &sys.AnnouncementRes{
		Announcement: dict.V,
	}, nil
}
