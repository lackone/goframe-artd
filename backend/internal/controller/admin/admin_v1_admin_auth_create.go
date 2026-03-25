package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminAuthCreate(ctx context.Context, req *v1.AdminAuthCreateReq) (res *v1.AdminAuthCreateRes, err error) {
	return c.aal.AdminAuthCreate(ctx, req)
}
