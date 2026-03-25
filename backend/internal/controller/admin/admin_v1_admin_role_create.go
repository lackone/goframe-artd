package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminRoleCreate(ctx context.Context, req *v1.AdminRoleCreateReq) (res *v1.AdminRoleCreateRes, err error) {
	return c.arl.AdminRoleCreate(ctx, req)
}
