// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AdminAuths is the golang structure for table admin_auths.
type AdminAuths struct {
	Id            uint   `json:"id"              orm:"id"              description:"ID"`             // ID
	Name          string `json:"name"            orm:"name"            description:"权限标识"`           // 权限标识
	Title         string `json:"title"           orm:"title"           description:"菜单名称"`           // 菜单名称
	Path          string `json:"path"            orm:"path"            description:"路由地址"`           // 路由地址
	Component     string `json:"component"       orm:"component"       description:"组件路径"`           // 组件路径
	Type          int    `json:"type"            orm:"type"            description:"类型(1:菜单,2:按钮)"`  // 类型(1:菜单,2:按钮)
	IsEnable      int    `json:"is_enable"       orm:"is_enable"       description:"是否启用(-1:否,1:是)"` // 是否启用(-1:否,1:是)
	Keepalive     int    `json:"keepalive"       orm:"keepalive"       description:"页面缓存(-1:否,1:是)"` // 页面缓存(-1:否,1:是)
	IsHide        int    `json:"is_hide"         orm:"is_hide"         description:"隐藏菜单(-1:否,1:是)"` // 隐藏菜单(-1:否,1:是)
	IsHideTab     int    `json:"is_hide_tab"     orm:"is_hide_tab"     description:"标签隐藏(-1:否,1:是)"` // 标签隐藏(-1:否,1:是)
	Link          string `json:"link"            orm:"link"            description:"外部链接"`           // 外部链接
	IsIframe      int    `json:"is_iframe"       orm:"is_iframe"       description:"是否内嵌(-1:否,1:是)"` // 是否内嵌(-1:否,1:是)
	ShowBadge     int    `json:"show_badge"      orm:"show_badge"      description:"显示徽章(-1:否,1:是)"` // 显示徽章(-1:否,1:是)
	ShowTextBadge string `json:"show_text_badge" orm:"show_text_badge" description:"文本徽章"`           // 文本徽章
	FixedTab      int    `json:"fixed_tab"       orm:"fixed_tab"       description:"固定标签(-1:否,1:是)"` // 固定标签(-1:否,1:是)
	ActivePath    string `json:"active_path"     orm:"active_path"     description:"激活路径"`           // 激活路径
	IsFullPage    int    `json:"is_full_page"    orm:"is_full_page"    description:"全屏页面(-1:否,1:是)"` // 全屏页面(-1:否,1:是)
	Pid           int    `json:"pid"             orm:"pid"             description:"父级"`             // 父级
	Icon          string `json:"icon"            orm:"icon"            description:"图标"`             // 图标
	Sort          int    `json:"sort"            orm:"sort"            description:"排序(越小越前)"`       // 排序(越小越前)
	Created       int    `json:"created"         orm:"created"         description:"创建时间"`           // 创建时间
	Updated       int    `json:"updated"         orm:"updated"         description:"更新时间"`           // 更新时间
	Deleted       int    `json:"deleted"         orm:"deleted"         description:"删除时间"`           // 删除时间
}
