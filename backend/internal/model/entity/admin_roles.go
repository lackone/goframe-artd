// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AdminRoles is the golang structure for table admin_roles.
type AdminRoles struct {
	Id      uint   `json:"id"       orm:"id"       description:"ID"`             // ID
	Name    string `json:"name"     orm:"name"     description:"角色名"`            // 角色名
	AuthIds string `json:"auth_ids" orm:"auth_ids" description:"权限ID(逗号分割)"`     // 权限ID(逗号分割)
	Status  int    `json:"status"   orm:"status"   description:"状态(-1:禁用,1:启用)"` // 状态(-1:禁用,1:启用)
	Remark  string `json:"remark"   orm:"remark"   description:"备注"`             // 备注
	Created int    `json:"created"  orm:"created"  description:"创建时间"`           // 创建时间
	Updated int    `json:"updated"  orm:"updated"  description:"更新时间"`           // 更新时间
	Deleted int    `json:"deleted"  orm:"deleted"  description:"删除时间"`           // 删除时间
}
