package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminSave(ctx context.Context, req *v1.AdminSaveReq) (res *v1.AdminSaveRes, err error) {
	return c.al.AdminSave(ctx, req)
}
