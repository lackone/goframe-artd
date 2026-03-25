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
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type AdminAuthLogic struct {
}

func NewAdminAuthLogic() *AdminAuthLogic {
	return &AdminAuthLogic{}
}

// AdminAuthSave 保存权限
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (al *AdminAuthLogic) AdminAuthSave(ctx context.Context, req *v1.AdminAuthSaveReq) (res *v1.AdminAuthSaveRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	data := r.GetMap()
	delete(data, "id")

	if req.Id > 0 {
		_, err = dao.AdminAuths.Ctx(ctx).Where("id = ?", req.Id).Update(data)
	} else {
		_, err = dao.AdminAuths.Ctx(ctx).Insert(data)
	}
	if err != nil {
		return nil, err
	}

	return &map[string]any{}, nil
}

// AdminAuthCreate 创建权限
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (al *AdminAuthLogic) AdminAuthCreate(ctx context.Context, req *v1.AdminAuthCreateReq) (res *v1.AdminAuthCreateRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	data := r.GetMap()
	delete(data, "id")

	if req.Id > 0 {
		_, err = dao.AdminAuths.Ctx(ctx).Where("id = ?", req.Id).Update(data)
	} else {
		_, err = dao.AdminAuths.Ctx(ctx).Insert(data)
	}
	if err != nil {
		return nil, err
	}

	return &map[string]any{}, nil
}

// AdminAuthMenuList 获取权限菜单列表
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (al *AdminAuthLogic) AdminAuthMenuList(ctx context.Context, req *v1.AdminAuthMenuListReq) (*v1.AdminAuthMenuListRes, error) {
	admin, err := utils.GetAdmin(ctx)
	if err != nil {
		return nil, err
	}

	authList, err := al.getAuthListByAdmin(ctx, admin, 0, 0)
	if err != nil {
		return nil, err
	}

	tree := al.MakeMenuTree(authList)
	menuList := TransformMenu(tree)

	return &menuList, nil
}

// MakeMenuTree 生成菜单树
// @param authList 权限列表
// @return 菜单树
func (al *AdminAuthLogic) MakeMenuTree(authList []map[string]any) []map[string]any {
	if len(authList) == 0 {
		return nil
	}

	var tree []map[string]any
	for _, auth := range authList {
		tree = append(tree, map[string]any{
			"id":        auth["id"],
			"path":      auth["path"],
			"name":      auth["name"],
			"component": auth["component"],
			"pid":       auth["pid"],
			"meta": map[string]any{
				"title":         auth["title"],
				"icon":          auth["icon"],
				"showBadge":     gconv.Int(auth["show_badge"]) == 1,
				"showTextBadge": auth["show_text_badge"],
				"isHide":        gconv.Int(auth["is_hide"]) == 1,
				"isHideTab":     gconv.Int(auth["is_hide_tab"]) == 1,
				"link":          auth["link"],
				"isIframe":      gconv.Int(auth["is_iframe"]) == 1,
				"keepAlive":     gconv.Int(auth["keepalive"]) == 1,
				"fixedTab":      gconv.Int(auth["fixed_tab"]) == 1,
				"isFullPage":    gconv.Int(auth["is_full_page"]) == 1,
				"activePath":    auth["active_path"],
				"isEnable":      gconv.Int(auth["is_enable"]) == 1,
				"created":       auth["created"],
				"updated":       auth["updated"],
				"type":          auth["type"],
				"pid":           auth["pid"],
				"sort":          auth["sort"],
			},
		})
	}
	return utils.MakeTree(tree, "id", "pid", "children")
}

// AdminAuthList 获取权限列表
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (al *AdminAuthLogic) AdminAuthList(ctx context.Context, req *v1.AdminAuthListReq) (*v1.AdminAuthListRes, error) {
	query := dao.AdminAuths.Ctx(ctx)

	// 构建查询条件
	if req.Title != "" {
		query = query.Where("title like ?", req.Title+"%")
	}
	if req.Path != "" {
		query = query.Where("path like ?", req.Path+"%")
	}
	if req.Component != "" {
		query = query.Where("component like ?", req.Component+"%")
	}

	// 设置排序
	query = query.Order("sort asc, id asc")

	ret, err := query.All()
	if err != nil {
		return nil, err
	}

	tree := al.MakeMenuTree(ret.List())
	menuList := TransformMenu(tree)

	return &menuList, nil
}

// AdminAuthDelete 删除权限
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (al *AdminAuthLogic) AdminAuthDelete(ctx context.Context, req *v1.AdminAuthDeleteReq) (*v1.AdminAuthDeleteRes, error) {
	ids := strings.Split(req.Ids, ",")
	if len(ids) == 0 {
		return nil, gerror.New("权限ID列表不能为空")
	}
	dao.AdminAuths.Ctx(ctx).Where("id in (?)", ids).Delete()
	return &map[string]any{}, nil
}

// getRoleIds 获取管理员角色ID列表
// @param ctx 上下文
// @param adminId 管理员ID
// @return 角色ID列表
// @return error 错误
func (al *AdminAuthLogic) getRoleIds(ctx context.Context, adminId int) ([]int, error) {
	ret, err := dao.AdminRoleAssocs.Ctx(ctx).Fields("role_id").
		Where("admin_id = ?", adminId).Array()
	if err != nil {
		return nil, err
	}
	return gconv.Ints(ret), nil
}

// getAuthIds 获取管理员权限ID列表
// @param ctx 上下文
// @param roleIds 角色ID列表
// @return 权限ID列表
// @return error 错误
func (al *AdminAuthLogic) getAuthIds(ctx context.Context, roleIds []int) ([]int, error) {
	ret, err := dao.AdminRoles.Ctx(ctx).Fields("auth_ids").
		Where("id in (?)", roleIds).
		Where("status = ?", STATUS_ENABLE).Array()

	if err != nil {
		return nil, err
	}

	var result []int
	seen := make(map[int]bool)

	for _, item := range ret {
		authIds := item.String()
		if authIds == "" {
			continue
		}
		ids := strings.Split(authIds, ",")
		for _, id := range ids {
			id = strings.TrimSpace(id)
			if id == "" {
				continue
			}
			intId := gconv.Int(id)
			if !seen[intId] {
				seen[intId] = true
				result = append(result, intId)
			}
		}
	}
	return result, nil
}

// getAuthList 获取权限列表
// @param ctx 上下文
// @param authIds 权限ID列表
// @param status 状态
// @param authType 权限类型
// @return 权限权限列表
// @return error 错误
func (al *AdminAuthLogic) getAuthList(ctx context.Context, authIds []int, status int, authType int) (gdb.List, error) {
	query := dao.AdminAuths.Ctx(ctx)

	// 构建查询条件
	query = query.Where("id in (?)", authIds)

	if status != 0 {
		query = query.Where("is_enable=?", status)
	}
	if authType != 0 {
		query = query.Where("type=?", authType)
	}

	query = query.Order("sort asc, id asc")

	ret, err := query.All()
	if err != nil {
		return nil, err
	}

	return ret.List(), nil
}

// getAuthListByAdmin 获取管理员权限列表
// @param ctx 上下文
// @param admin 管理员
// @param status 状态
// @param authType 权限类型
// @return 权限列表
// @return error 错误
func (al *AdminAuthLogic) getAuthListByAdmin(ctx context.Context, admin *entity.Admins, status int, authType int) (gdb.List, error) {
	if admin.IsSuper == IS_SUPER_YES {
		return al.GetAllAuthList(ctx, status, authType, nil)
	} else {
		roleIds, _ := al.getRoleIds(ctx, int(admin.Id))
		authIds, _ := al.getAuthIds(ctx, roleIds)
		return al.getAuthList(ctx, authIds, status, authType)
	}
}

// GetAllAuthList 获取所有权限列表
// @param ctx 上下文
// @param status 状态
// @param authType 权限类型
// @param fields 字段
// @return 权限列表
// @return error 错误
func (al *AdminAuthLogic) GetAllAuthList(ctx context.Context, status int, authType int, fields []any) (gdb.List, error) {
	query := dao.AdminAuths.Ctx(ctx)

	// 构建查询条件
	if status != 0 {
		query = query.Where("is_enable=?", status)
	}
	if authType != 0 {
		query = query.Where("type=?", authType)
	}

	// 设置排序
	query = query.Order("sort asc, id asc")

	// 执行查询
	if len(fields) > 0 {
		ret, err := query.Fields(fields...).All()
		if err != nil {
			return nil, err
		}
		return ret.List(), nil
	} else {
		ret, err := query.All()
		if err != nil {
			return nil, err
		}
		return ret.List(), nil
	}
}

// AuthList 获取权限列表
// @param ctx 上下文
// @return 权限列表
// @return error 错误
func (al *AdminAuthLogic) AuthList(ctx context.Context) (gdb.List, error) {
	return al.GetAllAuthList(ctx, STATUS_ENABLE, 0, []any{"id as value", "title as label", "pid"})
}

// TransformMenu 转换菜单，提取按钮权限
// @param menuList 菜单列表
// @return 转换后的菜单列表
func TransformMenu(menuList []map[string]any) []map[string]any {
	if len(menuList) == 0 {
		return nil
	}

	var result []map[string]any

	for _, item := range menuList {
		newItem := gconv.Map(item)

		children, ok := newItem["children"].([]map[string]any)
		if ok && len(children) > 0 {
			var authList []map[string]any
			var submenuChildren []map[string]any

			for _, child := range children {
				childMap := gconv.Map(child)
				metaMap, ok := childMap["meta"].(map[string]any)
				if !ok {
					metaMap = make(map[string]any)
				}

				authType := gconv.Int(metaMap["type"])

				if authType == AUTH_TYPE_BTN {
					authList = append(authList, map[string]any{
						"id":       childMap["id"],
						"pid":      metaMap["pid"],
						"sort":     metaMap["sort"],
						"title":    metaMap["title"],
						"authMark": childMap["name"],
						"created":  metaMap["created"],
						"updated":  metaMap["updated"],
						"type":     authType,
						"isEnable": gconv.Int(metaMap["isEnable"]) == 1,
					})
				} else {
					submenuChildren = append(submenuChildren, TransformMenu([]map[string]any{childMap})[0])
				}
			}

			newItemMeta, ok := newItem["meta"].(map[string]any)
			if !ok {
				newItemMeta = make(map[string]any)
			}
			if authList == nil {
				authList = []map[string]any{}
			}
			newItemMeta["authList"] = authList
			newItem["meta"] = newItemMeta

			if children, ok := newItem["children"].([]map[string]any); ok && len(children) > 0 {
				newItem["children"] = submenuChildren
			} else {
				newItem["children"] = []map[string]any{}
			}
		}

		result = append(result, newItem)
	}

	return result
}
