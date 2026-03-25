package v1

import (
	"backend/api/common"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminRoleListReq struct {
	g.Meta `path:"/role/list" method:"get" summary:"获取角色列表"`
	common.PageReq
	Name string `json:"name" dc:"角色名称"`
	StartTime string `json:"start_time" dc:"开始时间"`
	EndTime string `json:"end_time" dc:"结束时间"`
	Status int `json:"status" dc:"角色状态"`
}

type AdminRoleListRes struct {
	common.PageRes
}

type AdminRoleCreateReq struct {
	g.Meta `path:"/role/save" method:"post" summary:"创建角色"`
	Id     int `json:"id" dc:"角色ID"`
}

type AdminRoleCreateRes = map[string]any

type AdminRoleSaveReq struct {
	g.Meta `path:"/role/save/:id" method:"post" summary:"保存角色"`
	Id     int `json:"id" dc:"角色ID"`
}

type AdminRoleSaveRes = map[string]any

type AdminRoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"post" summary:"删除角色"`
	Ids    string `json:"ids" v:"required" dc:"角色ID列表"`
}

type AdminRoleDeleteRes = map[string]any
