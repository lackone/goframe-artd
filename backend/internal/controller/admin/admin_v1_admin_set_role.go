package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminSetRole(ctx context.Context, req *v1.AdminSetRoleReq) (res *v1.AdminSetRoleRes, err error) {
	return c.al.AdminSetRole(ctx, req)
}
