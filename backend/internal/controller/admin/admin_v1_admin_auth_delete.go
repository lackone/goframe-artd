package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminAuthDelete(ctx context.Context, req *v1.AdminAuthDeleteReq) (res *v1.AdminAuthDeleteRes, err error) {
	return c.aal.AdminAuthDelete(ctx, req)
}
