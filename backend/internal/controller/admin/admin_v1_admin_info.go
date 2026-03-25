package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminInfo(ctx context.Context, req *v1.AdminInfoReq) (res *v1.AdminInfoRes, err error) {
	return c.al.AdminInfo(ctx, req)
}
