package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminRoleList(ctx context.Context, req *v1.AdminRoleListReq) (res *v1.AdminRoleListRes, err error) {
	return c.arl.AdminRoleList(ctx, req)
}
