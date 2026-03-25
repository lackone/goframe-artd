package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminAuthList(ctx context.Context, req *v1.AdminAuthListReq) (res *v1.AdminAuthListRes, err error) {
	return c.aal.AdminAuthList(ctx, req)
}
