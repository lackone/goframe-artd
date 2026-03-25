package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	roleList, err := c.arl.RoleList(ctx)
	if err != nil {
		return nil, err
	}
	return &roleList, nil
}
