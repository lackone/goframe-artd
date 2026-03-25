package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "backend/api/admin/v1"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "file is required")
	}
	result, err := c.fs.Upload(req.File)
	if err != nil {
		return nil, err
	}
	return &v1.UploadRes{
		Ext:  result["ext"].(string),
		Path: result["path"].(string),
		Url:  result["url"].(string),
		Name: result["name"].(string),
		Size: result["size"].(int64),
	}, nil
}
