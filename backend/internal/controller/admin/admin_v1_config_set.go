package admin

import (
	"context"

	"backend/api/admin/v1"
)

func (c *ControllerV1) ConfigSet(ctx context.Context, req *v1.ConfigSetReq) (res *v1.ConfigSetRes, err error) {
	return c.cl.ConfigSet(ctx, req)
}
