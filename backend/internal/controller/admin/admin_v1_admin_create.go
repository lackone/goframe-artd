package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminCreate(ctx context.Context, req *v1.AdminCreateReq) (res *v1.AdminCreateRes, err error) {
	return c.al.AdminCreate(ctx, req)
}
