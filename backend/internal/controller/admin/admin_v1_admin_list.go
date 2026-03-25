package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminList(ctx context.Context, req *v1.AdminListReq) (res *v1.AdminListRes, err error) {
	return c.al.AdminList(ctx, req)
}
