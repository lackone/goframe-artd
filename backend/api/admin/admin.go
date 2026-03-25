// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"backend/api/admin/v1"
)

type IAdminV1 interface {
	AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error)
	AdminInfo(ctx context.Context, req *v1.AdminInfoReq) (res *v1.AdminInfoRes, err error)
	AdminList(ctx context.Context, req *v1.AdminListReq) (res *v1.AdminListRes, err error)
	AdminDelete(ctx context.Context, req *v1.AdminDeleteReq) (res *v1.AdminDeleteRes, err error)
	AdminCreate(ctx context.Context, req *v1.AdminCreateReq) (res *v1.AdminCreateRes, err error)
	AdminSave(ctx context.Context, req *v1.AdminSaveReq) (res *v1.AdminSaveRes, err error)
	AdminSetRole(ctx context.Context, req *v1.AdminSetRoleReq) (res *v1.AdminSetRoleRes, err error)
	AdminAuthList(ctx context.Context, req *v1.AdminAuthListReq) (res *v1.AdminAuthListRes, err error)
	AdminAuthMenuList(ctx context.Context, req *v1.AdminAuthMenuListReq) (res *v1.AdminAuthMenuListRes, err error)
	AdminAuthCreate(ctx context.Context, req *v1.AdminAuthCreateReq) (res *v1.AdminAuthCreateRes, err error)
	AdminAuthSave(ctx context.Context, req *v1.AdminAuthSaveReq) (res *v1.AdminAuthSaveRes, err error)
	AdminAuthDelete(ctx context.Context, req *v1.AdminAuthDeleteReq) (res *v1.AdminAuthDeleteRes, err error)
	AdminRoleList(ctx context.Context, req *v1.AdminRoleListReq) (res *v1.AdminRoleListRes, err error)
	AdminRoleCreate(ctx context.Context, req *v1.AdminRoleCreateReq) (res *v1.AdminRoleCreateRes, err error)
	AdminRoleSave(ctx context.Context, req *v1.AdminRoleSaveReq) (res *v1.AdminRoleSaveRes, err error)
	AdminRoleDelete(ctx context.Context, req *v1.AdminRoleDeleteReq) (res *v1.AdminRoleDeleteRes, err error)
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
	RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error)
	AuthList(ctx context.Context, req *v1.AuthListReq) (res *v1.AuthListRes, err error)
	ConfigSet(ctx context.Context, req *v1.ConfigSetReq) (res *v1.ConfigSetRes, err error)
	ConfigGet(ctx context.Context, req *v1.ConfigGetReq) (res *v1.ConfigGetRes, err error)
}
