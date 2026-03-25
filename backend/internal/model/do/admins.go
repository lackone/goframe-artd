// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Admins is the golang structure of table tb_admins for DAO operations like Where/Data.
type Admins struct {
	g.Meta        `orm:"table:tb_admins, do:true"`
	Id            any // ID
	Account       any // 账号
	NickName      any // 昵称
	RealName      any // 真实姓名
	Avatar        any // 头像
	Sex           any // 性别(0:未知,1:男,2:女)
	Salt          any // 加密盐
	Password      any // 密码(md5(md5(salt) . password))
	Phone         any // 手机号
	Tel           any // 座机
	Email         any // 邮箱
	Weixin        any // 微信
	LastLoginIp   any // 最后登录IP
	LastLoginTime any // 最后登录时间
	IsSuper       any // 超级管理员(-1:否,1:是)
	Status        any // 状态(-1:禁用,1:启用)
	Address       any // 地址
	Remark        any // 备注
	Created       any // 创建时间
	Updated       any // 更新时间
	Deleted       any // 删除时间
}
