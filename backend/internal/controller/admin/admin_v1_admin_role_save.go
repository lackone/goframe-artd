package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) AdminRoleSave(ctx context.Context, req *v1.AdminRoleSaveReq) (res *v1.AdminRoleSaveRes, err error) {
	return c.arl.AdminRoleSave(ctx, req)
}
