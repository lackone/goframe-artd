package v1

import (
	"backend/api/common"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminLoginReq struct {
	g.Meta   `path:"/admin/login" method:"post" summary:"登录"`
	Account  string `json:"account" v:"required" dc:"账号" `
	Password string `json:"password" v:"required" dc:"密码" `
}

type AdminLoginRes struct {
	Token string `json:"token" dc:"登录token"`
}

type AdminInfoReq struct {
	g.Meta `path:"/admin/info" method:"get" summary:"获取用户信息"`
}

type AdminInfoRes = map[string]any

type AdminListReq struct {
	g.Meta `path:"/admin/list" method:"get" summary:"获取用户列表"`
	common.PageReq
	Account   string `json:"account" dc:"账号"`
	Phone     string `json:"phone" dc:"手机号"`
	Email     string `json:"email" dc:"邮箱"`
	StartTime string `json:"start_time" dc:"开始时间"`
	EndTime   string `json:"end_time" dc:"结束时间"`
	Status    int    `json:"status" dc:"状态"`
}

type AdminListRes struct {
	common.PageRes
}

type AdminDeleteReq struct {
	g.Meta `path:"/admin/delete" method:"post" summary:"删除用户"`
	Ids    string `json:"ids" v:"required" dc:"用户ID列表"`
}

type AdminDeleteRes = map[string]any

type AdminCreateReq struct {
	g.Meta `path:"/admin/save" method:"post" summary:"创建用户"`
	Id     int `json:"id" dc:"用户ID"`
}

type AdminCreateRes = map[string]any

type AdminSaveReq struct {
	g.Meta `path:"/admin/save/:id" method:"post" summary:"保存用户"`
	Id     int `json:"id" dc:"用户ID"`
}

type AdminSaveRes = map[string]any

type AdminSetRoleReq struct {
	g.Meta `path:"/admin/set_role/:id" method:"post" summary:"设置用户角色"`
	Id     int `json:"id" dc:"用户ID"`
	RoleIdIds string `json:"role_ids" v:"required" dc:"角色ID列表"`
}

type AdminSetRoleRes = map[string]any
