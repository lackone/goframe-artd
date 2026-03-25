package admin

import (
	"context"

	v1 "backend/api/admin/v1"
)

func (c *ControllerV1) AdminAuthMenuList(ctx context.Context, req *v1.AdminAuthMenuListReq) (res *v1.AdminAuthMenuListRes, err error) {
	return c.aal.AdminAuthMenuList(ctx, req)
}
