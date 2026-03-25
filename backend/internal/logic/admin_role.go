package logic

import (
	v1 "backend/api/admin/v1"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/internal/utils"
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type AdminRoleLogic struct {
}

func NewAdminRoleLogic() *AdminRoleLogic {
	return &AdminRoleLogic{}
}

type Roles struct {
    g.Meta `orm:"table:admin_roles"`
    Id        int    `json:"id"`
    Name    string `json:"name"`
}

// Save 保存角色
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (rl *AdminRoleLogic) AdminRoleSave(ctx context.Context, req *v1.AdminRoleSaveReq) (res *v1.AdminRoleSaveRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	data := r.GetMap()
	delete(data, "id")

	if req.Id > 0 {
		_, err = dao.AdminRoles.Ctx(ctx).Where("id = ?", req.Id).Update(data)
	} else {
		_, err = dao.AdminRoles.Ctx(ctx).Insert(data)
	}
	if err != nil {
		return nil, err
	}

	return &map[string]any{}, nil
}

// Create 创建角色
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (rl *AdminRoleLogic) AdminRoleCreate(ctx context.Context, req *v1.AdminRoleCreateReq) (res *v1.AdminRoleCreateRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	data := r.GetMap()
	delete(data, "id")

	if req.Id > 0 {
		_, err = dao.AdminRoles.Ctx(ctx).Where("id = ?", req.Id).Update(data)
	} else {
		_, err = dao.AdminRoles.Ctx(ctx).Insert(data)
	}
	if err != nil {
		return nil, err
	}

	return &map[string]any{}, nil
}

// List 获取角色列表
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (rl *AdminRoleLogic) AdminRoleList(ctx context.Context, req *v1.AdminRoleListReq) (*v1.AdminRoleListRes, error) {
	query := dao.AdminRoles.Ctx(ctx)

	// 构建查询条件
	if req.Name != "" {
		query = query.Where("name like ?", req.Name+"%")
	}
	if req.StartTime != "" {
		query = query.Where("created >= ?", utils.ToTimestamp(req.StartTime))
	}
	if req.EndTime != "" {
		query = query.Where("created <= ?", utils.ToTimestamp(req.EndTime)+86400)
	}
	if req.Status != 0 {
		query = query.Where("status=?", req.Status)
	}

	query = query.Order("id desc")

	// 分页查询
	pageRes, err := utils.ToList[entity.AdminRoles](query, &req.PageReq)
	if err != nil {
		return nil, err
	}

	return &v1.AdminRoleListRes{
		PageRes: *pageRes,
	}, nil
}

// Delete 删除角色
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (rl *AdminRoleLogic) AdminRoleDelete(ctx context.Context, req *v1.AdminRoleDeleteReq) (*v1.AdminRoleDeleteRes, error) {
	ids := strings.Split(req.Ids, ",")
	if len(ids) == 0 {
		return nil, gerror.New("角色ID列表不能为空")
	}
	dao.AdminRoles.Ctx(ctx).Where("id in (?)", ids).Delete()
	return &map[string]any{}, nil
}

// RoleList 获取角色列表
// @param ctx 上下文
// @return 角色列表
// @return error 错误
func (rl *AdminRoleLogic) RoleList(ctx context.Context) (gdb.List, error) {
	query := dao.AdminRoles.Ctx(ctx)
	ret, err := query.Fields("id, name").Where("status=?", STATUS_ENABLE).All()
	if err != nil {
		return nil, err
	}
	return ret.List(), nil
}
