package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadReq struct {
	g.Meta `path:"/common/upload" method:"post" mime:"multipart/form-data" summary:"文件上传"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type UploadRes struct {
	Ext  string `json:"ext" dc:"文件扩展名"`
	Path string `json:"path" dc:"文件路径"`
	Url  string `json:"url" dc:"文件URL"`
	Name string `json:"name" dc:"文件名"`
	Size int64  `json:"size" dc:"文件大小"`
}

type RoleListReq struct {
	g.Meta `path:"/common/role_list" method:"get" summary:"角色列表"`
}

type RoleListRes = []map[string]any

type AuthListReq struct {
	g.Meta `path:"/common/auth_list" method:"get" summary:"权限列表"`
}

type AuthListRes = []map[string]any