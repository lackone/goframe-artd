// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminAuths is the golang structure of table tb_admin_auths for DAO operations like Where/Data.
type AdminAuths struct {
	g.Meta        `orm:"table:tb_admin_auths, do:true"`
	Id            any // ID
	Name          any // 权限标识
	Title         any // 菜单名称
	Path          any // 路由地址
	Component     any // 组件路径
	Type          any // 类型(1:菜单,2:按钮)
	IsEnable      any // 是否启用(-1:否,1:是)
	Keepalive     any // 页面缓存(-1:否,1:是)
	IsHide        any // 隐藏菜单(-1:否,1:是)
	IsHideTab     any // 标签隐藏(-1:否,1:是)
	Link          any // 外部链接
	IsIframe      any // 是否内嵌(-1:否,1:是)
	ShowBadge     any // 显示徽章(-1:否,1:是)
	ShowTextBadge any // 文本徽章
	FixedTab      any // 固定标签(-1:否,1:是)
	ActivePath    any // 激活路径
	IsFullPage    any // 全屏页面(-1:否,1:是)
	Pid           any // 父级
	Icon          any // 图标
	Sort          any // 排序(越小越前)
	Created       any // 创建时间
	Updated       any // 更新时间
	Deleted       any // 删除时间
}
