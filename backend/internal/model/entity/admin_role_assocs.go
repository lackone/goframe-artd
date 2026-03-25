// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AdminRoleAssocs is the golang structure for table admin_role_assocs.
type AdminRoleAssocs struct {
	Id      uint `json:"id"       orm:"id"       description:"ID"`   // ID
	AdminId int  `json:"admin_id" orm:"admin_id" description:"用户ID"` // 用户ID
	RoleId  int  `json:"role_id"  orm:"role_id"  description:"角色ID"` // 角色ID
	Created int  `json:"created"  orm:"created"  description:"创建时间"` // 创建时间
	Updated int  `json:"updated"  orm:"updated"  description:"更新时间"` // 更新时间
	Deleted int  `json:"deleted"  orm:"deleted"  description:"删除时间"` // 删除时间
}
