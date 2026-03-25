// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoles is the golang structure of table tb_admin_roles for DAO operations like Where/Data.
type AdminRoles struct {
	g.Meta  `orm:"table:tb_admin_roles, do:true"`
	Id      any // ID
	Name    any // 角色名
	AuthIds any // 权限ID(逗号分割)
	Status  any // 状态(-1:禁用,1:启用)
	Remark  any // 备注
	Created any // 创建时间
	Updated any // 更新时间
	Deleted any // 删除时间
}
