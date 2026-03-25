// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleAssocs is the golang structure of table tb_admin_role_assocs for DAO operations like Where/Data.
type AdminRoleAssocs struct {
	g.Meta  `orm:"table:tb_admin_role_assocs, do:true"`
	Id      any // ID
	AdminId any // 用户ID
	RoleId  any // 角色ID
	Created any // 创建时间
	Updated any // 更新时间
	Deleted any // 删除时间
}
