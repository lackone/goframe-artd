package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminDelete(ctx context.Context, req *v1.AdminDeleteReq) (res *v1.AdminDeleteRes, err error) {
	return c.al.AdminDelete(ctx, req)
}
