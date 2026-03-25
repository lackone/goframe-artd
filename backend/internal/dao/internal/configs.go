// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigsDao is the data access object for the table tb_configs.
type ConfigsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ConfigsColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ConfigsColumns defines and stores column names for the table tb_configs.
type ConfigsColumns struct {
	Id      string // ID
	Name    string // 中文名
	Key     string // 键
	SubKey  string // 子键
	Value   string // 内容
	Pid     string // 父级
	Sort    string // 排序(越小越前)
	AdminId string // 操作人
	Created string // 创建时间
	Updated string // 更新时间
	Deleted string // 删除时间
}

// configsColumns holds the columns for the table tb_configs.
var configsColumns = ConfigsColumns{
	Id:      "id",
	Name:    "name",
	Key:     "key",
	SubKey:  "sub_key",
	Value:   "value",
	Pid:     "pid",
	Sort:    "sort",
	AdminId: "admin_id",
	Created: "created",
	Updated: "updated",
	Deleted: "deleted",
}

// NewConfigsDao creates and returns a new DAO object for table data access.
func NewConfigsDao(handlers ...gdb.ModelHandler) *ConfigsDao {
	return &ConfigsDao{
		group:    "default",
		table:    "tb_configs",
		columns:  configsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ConfigsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ConfigsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ConfigsDao) Columns() ConfigsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ConfigsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ConfigsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ConfigsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
