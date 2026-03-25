// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleAssocsDao is the data access object for the table tb_admin_role_assocs.
type AdminRoleAssocsDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  AdminRoleAssocsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// AdminRoleAssocsColumns defines and stores column names for the table tb_admin_role_assocs.
type AdminRoleAssocsColumns struct {
	Id      string // ID
	AdminId string // 用户ID
	RoleId  string // 角色ID
	Created string // 创建时间
	Updated string // 更新时间
	Deleted string // 删除时间
}

// adminRoleAssocsColumns holds the columns for the table tb_admin_role_assocs.
var adminRoleAssocsColumns = AdminRoleAssocsColumns{
	Id:      "id",
	AdminId: "admin_id",
	RoleId:  "role_id",
	Created: "created",
	Updated: "updated",
	Deleted: "deleted",
}

// NewAdminRoleAssocsDao creates and returns a new DAO object for table data access.
func NewAdminRoleAssocsDao(handlers ...gdb.ModelHandler) *AdminRoleAssocsDao {
	return &AdminRoleAssocsDao{
		group:    "default",
		table:    "tb_admin_role_assocs",
		columns:  adminRoleAssocsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminRoleAssocsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminRoleAssocsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminRoleAssocsDao) Columns() AdminRoleAssocsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminRoleAssocsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminRoleAssocsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminRoleAssocsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
