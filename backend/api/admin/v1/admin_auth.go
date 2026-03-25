package v1

import (
	"backend/api/common"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminAuthListReq struct {
	g.Meta `path:"/auth/list" method:"get" summary:"获取权限列表"`
	common.PageReq
	Title string `json:"title" dc:"权限标题"`
	Path  string `json:"path" dc:"权限路径"`
	Component string `json:"component" dc:"权限组件"`
}

type AdminAuthListRes = []map[string]any

type AdminAuthMenuListReq struct {
	g.Meta `path:"/auth/menu_list" method:"get" summary:"获取权限菜单列表"`
}

type AdminAuthMenuListRes  = []map[string]any

type AdminAuthCreateReq struct {
	g.Meta `path:"/auth/save" method:"post" summary:"创建权限"`
	Id     int `json:"id" dc:"权限ID"`
}

type AdminAuthCreateRes = map[string]any

type AdminAuthSaveReq struct {
	g.Meta `path:"/auth/save/:id" method:"post" summary:"保存权限"`
	Id     int `json:"id" dc:"权限ID"`
}

type AdminAuthSaveRes = map[string]any

type AdminAuthDeleteReq struct {
	g.Meta `path:"/auth/delete" method:"post" summary:"删除权限"`
	Ids    string `json:"ids" v:"required" dc:"权限ID列表"`
}

type AdminAuthDeleteRes = map[string]any