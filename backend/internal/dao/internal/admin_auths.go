// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminAuthsDao is the data access object for the table tb_admin_auths.
type AdminAuthsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminAuthsColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminAuthsColumns defines and stores column names for the table tb_admin_auths.
type AdminAuthsColumns struct {
	Id            string // ID
	Name          string // 权限标识
	Title         string // 菜单名称
	Path          string // 路由地址
	Component     string // 组件路径
	Type          string // 类型(1:菜单,2:按钮)
	IsEnable      string // 是否启用(-1:否,1:是)
	Keepalive     string // 页面缓存(-1:否,1:是)
	IsHide        string // 隐藏菜单(-1:否,1:是)
	IsHideTab     string // 标签隐藏(-1:否,1:是)
	Link          string // 外部链接
	IsIframe      string // 是否内嵌(-1:否,1:是)
	ShowBadge     string // 显示徽章(-1:否,1:是)
	ShowTextBadge string // 文本徽章
	FixedTab      string // 固定标签(-1:否,1:是)
	ActivePath    string // 激活路径
	IsFullPage    string // 全屏页面(-1:否,1:是)
	Pid           string // 父级
	Icon          string // 图标
	Sort          string // 排序(越小越前)
	Created       string // 创建时间
	Updated       string // 更新时间
	Deleted       string // 删除时间
}

// adminAuthsColumns holds the columns for the table tb_admin_auths.
var adminAuthsColumns = AdminAuthsColumns{
	Id:            "id",
	Name:          "name",
	Title:         "title",
	Path:          "path",
	Component:     "component",
	Type:          "type",
	IsEnable:      "is_enable",
	Keepalive:     "keepalive",
	IsHide:        "is_hide",
	IsHideTab:     "is_hide_tab",
	Link:          "link",
	IsIframe:      "is_iframe",
	ShowBadge:     "show_badge",
	ShowTextBadge: "show_text_badge",
	FixedTab:      "fixed_tab",
	ActivePath:    "active_path",
	IsFullPage:    "is_full_page",
	Pid:           "pid",
	Icon:          "icon",
	Sort:          "sort",
	Created:       "created",
	Updated:       "updated",
	Deleted:       "deleted",
}

// NewAdminAuthsDao creates and returns a new DAO object for table data access.
func NewAdminAuthsDao(handlers ...gdb.ModelHandler) *AdminAuthsDao {
	return &AdminAuthsDao{
		group:    "default",
		table:    "tb_admin_auths",
		columns:  adminAuthsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminAuthsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminAuthsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminAuthsDao) Columns() AdminAuthsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminAuthsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminAuthsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AdminAuthsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
