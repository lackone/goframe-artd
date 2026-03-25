package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminRoleDelete(ctx context.Context, req *v1.AdminRoleDeleteReq) (res *v1.AdminRoleDeleteRes, err error) {
	return c.arl.AdminRoleDelete(ctx, req)
}
