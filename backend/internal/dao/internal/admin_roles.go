// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRolesDao is the data access object for the table tb_admin_roles.
type AdminRolesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminRolesColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminRolesColumns defines and stores column names for the table tb_admin_roles.
type AdminRolesColumns struct {
	Id      string // ID
	Name    string // 角色名
	AuthIds string // 权限ID(逗号分割)
	Status  string // 状态(-1:禁用,1:启用)
	Remark  string // 备注
	Created string // 创建时间
	Updated string // 更新时间
	Deleted string // 删除时间
}

// adminRolesColumns holds the columns for the table tb_admin_roles.
var adminRolesColumns = AdminRolesColumns{
	Id:      "id",
	Name:    "name",
	AuthIds: "auth_ids",
	Status:  "status",
	Remark:  "remark",
	Created: "created",
	Updated: "updated",
	Deleted: "deleted",
}

// NewAdminRolesDao creates and returns a new DAO object for table data access.
func NewAdminRolesDao(handlers ...gdb.ModelHandler) *AdminRolesDao {
	return &AdminRolesDao{
		group:    "default",
		table:    "tb_admin_roles",
		columns:  adminRolesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminRolesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminRolesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminRolesDao) Columns() AdminRolesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminRolesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminRolesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminRolesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
