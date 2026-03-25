// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Configs is the golang structure of table tb_configs for DAO operations like Where/Data.
type Configs struct {
	g.Meta  `orm:"table:tb_configs, do:true"`
	Id      any // ID
	Name    any // 中文名
	Key     any // 键
	SubKey  any // 子键
	Value   any // 内容
	Pid     any // 父级
	Sort    any // 排序(越小越前)
	AdminId any // 操作人
	Created any // 创建时间
	Updated any // 更新时间
	Deleted any // 删除时间
}
