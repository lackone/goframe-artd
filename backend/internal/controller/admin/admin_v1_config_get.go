package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) ConfigGet(ctx context.Context, req *v1.ConfigGetReq) (res *v1.ConfigGetRes, err error) {
	return c.cl.ConfigGet(ctx, req)
}
