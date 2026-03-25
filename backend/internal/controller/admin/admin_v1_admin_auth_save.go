package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminAuthSave(ctx context.Context, req *v1.AdminAuthSaveReq) (res *v1.AdminAuthSaveRes, err error) {
	return c.aal.AdminAuthSave(ctx, req)
}
