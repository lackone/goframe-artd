// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Admins is the golang structure for table admins.
type Admins struct {
	Id            uint   `json:"id"              orm:"id"              description:"ID"`                            // ID
	Account       string `json:"account"         orm:"account"         description:"账号"`                            // 账号
	NickName      string `json:"nick_name"       orm:"nick_name"       description:"昵称"`                            // 昵称
	RealName      string `json:"real_name"       orm:"real_name"       description:"真实姓名"`                          // 真实姓名
	Avatar        string `json:"avatar"          orm:"avatar"          description:"头像"`                            // 头像
	Sex           int    `json:"sex"             orm:"sex"             description:"性别(0:未知,1:男,2:女)"`              // 性别(0:未知,1:男,2:女)
	Salt          string `json:"salt"            orm:"salt"            description:"加密盐"`                           // 加密盐
	Password      string `json:"password"        orm:"password"        description:"密码(md5(md5(salt) . password))"` // 密码(md5(md5(salt) . password))
	Phone         string `json:"phone"           orm:"phone"           description:"手机号"`                           // 手机号
	Tel           string `json:"tel"             orm:"tel"             description:"座机"`                            // 座机
	Email         string `json:"email"           orm:"email"           description:"邮箱"`                            // 邮箱
	Weixin        string `json:"weixin"          orm:"weixin"          description:"微信"`                            // 微信
	LastLoginIp   string `json:"last_login_ip"   orm:"last_login_ip"   description:"最后登录IP"`                        // 最后登录IP
	LastLoginTime int    `json:"last_login_time" orm:"last_login_time" description:"最后登录时间"`                        // 最后登录时间
	IsSuper       int    `json:"is_super"        orm:"is_super"        description:"超级管理员(-1:否,1:是)"`               // 超级管理员(-1:否,1:是)
	Status        int    `json:"status"          orm:"status"          description:"状态(-1:禁用,1:启用)"`                // 状态(-1:禁用,1:启用)
	Address       string `json:"address"         orm:"address"         description:"地址"`                            // 地址
	Remark        string `json:"remark"          orm:"remark"          description:"备注"`                            // 备注
	Created       int    `json:"created"         orm:"created"         description:"创建时间"`                          // 创建时间
	Updated       int    `json:"updated"         orm:"updated"         description:"更新时间"`                          // 更新时间
	Deleted       int    `json:"deleted"         orm:"deleted"         description:"删除时间"`                          // 删除时间
}
