package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	return c.al.AdminLogin(ctx, req)
}
