// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminsDao is the data access object for the table tb_admins.
type AdminsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminsColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminsColumns defines and stores column names for the table tb_admins.
type AdminsColumns struct {
	Id            string // ID
	Account       string // 账号
	NickName      string // 昵称
	RealName      string // 真实姓名
	Avatar        string // 头像
	Sex           string // 性别(0:未知,1:男,2:女)
	Salt          string // 加密盐
	Password      string // 密码(md5(md5(salt) . password))
	Phone         string // 手机号
	Tel           string // 座机
	Email         string // 邮箱
	Weixin        string // 微信
	LastLoginIp   string // 最后登录IP
	LastLoginTime string // 最后登录时间
	IsSuper       string // 超级管理员(-1:否,1:是)
	Status        string // 状态(-1:禁用,1:启用)
	Address       string // 地址
	Remark        string // 备注
	Created       string // 创建时间
	Updated       string // 更新时间
	Deleted       string // 删除时间
}

// adminsColumns holds the columns for the table tb_admins.
var adminsColumns = AdminsColumns{
	Id:            "id",
	Account:       "account",
	NickName:      "nick_name",
	RealName:      "real_name",
	Avatar:        "avatar",
	Sex:           "sex",
	Salt:          "salt",
	Password:      "password",
	Phone:         "phone",
	Tel:           "tel",
	Email:         "email",
	Weixin:        "weixin",
	LastLoginIp:   "last_login_ip",
	LastLoginTime: "last_login_time",
	IsSuper:       "is_super",
	Status:        "status",
	Address:       "address",
	Remark:        "remark",
	Created:       "created",
	Updated:       "updated",
	Deleted:       "deleted",
}

// NewAdminsDao creates and returns a new DAO object for table data access.
func NewAdminsDao(handlers ...gdb.ModelHandler) *AdminsDao {
	return &AdminsDao{
		group:    "default",
		table:    "tb_admins",
		columns:  adminsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminsDao) Columns() AdminsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
