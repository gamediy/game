package user_svc

import (
	"context"
	"game/app/user/api/user/user"
)

func (userSvc) ListMailBox(ctx context.Context, req *user.ListMailBoxReq) (*user.ListMailBoxRes, error) {
	return userClient.ListMailBox(ctx, req)
}
