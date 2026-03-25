// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Configs is the golang structure for table configs.
type Configs struct {
	Id      uint   `json:"id"       orm:"id"       description:"ID"`       // ID
	Name    string `json:"name"     orm:"name"     description:"中文名"`      // 中文名
	Key     string `json:"key"      orm:"key"      description:"键"`        // 键
	SubKey  string `json:"sub_key"  orm:"sub_key"  description:"子键"`       // 子键
	Value   string `json:"value"    orm:"value"    description:"内容"`       // 内容
	Pid     int    `json:"pid"      orm:"pid"      description:"父级"`       // 父级
	Sort    int    `json:"sort"     orm:"sort"     description:"排序(越小越前)"` // 排序(越小越前)
	AdminId int    `json:"admin_id" orm:"admin_id" description:"操作人"`      // 操作人
	Created int    `json:"created"  orm:"created"  description:"创建时间"`     // 创建时间
	Updated int    `json:"updated"  orm:"updated"  description:"更新时间"`     // 更新时间
	Deleted int    `json:"deleted"  orm:"deleted"  description:"删除时间"`     // 删除时间
}
