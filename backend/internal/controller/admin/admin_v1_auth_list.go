package admin

import (
	"context"

	"backend/api/admin/v1"
	"backend/internal/utils"
)

func (c *ControllerV1) AuthList(ctx context.Context, req *v1.AuthListReq) (res *v1.AuthListRes, err error) {
	authList, err := c.aal.AuthList(ctx)
	if err != nil {
		return nil, err
	}
	tree := utils.MakeTree(authList, "value", "pid", "children")
	return &tree, nil
}
