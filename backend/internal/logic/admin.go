package logic

import (
	"context"
	"strings"
	"time"

	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/internal/utils"

	v1 "backend/api/admin/v1"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

type AdminLogic struct {
}

func NewAdminLogic() *AdminLogic {
	return &AdminLogic{}
}

type JwtClaims struct {
	Id      uint
	Account string
	jwt.RegisteredClaims
}

type ExtAdmins struct {
	entity.Admins
	Roles []map[string]any `json:"roles"`
}

// Login 登录
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (l *AdminLogic) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	var admin entity.Admins
	err = dao.Admins.Ctx(ctx).Where("account = ?", req.Account).Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin.Id == 0 {
		return nil, gerror.New("用户不存在")
	}
	if admin.Status == STATUS_DISABLE {
		return nil, gerror.New("用户已禁用")
	}
	pwd := l.makePassword(req.Password, admin.Salt)
	if pwd != admin.Password {
		return nil, gerror.New("密码错误")
	}

	token, err := l.createTokenWithTTL(&admin, consts.JwtExpire)
	if err != nil {
		return nil, err
	}

	return &v1.AdminLoginRes{
		Token: "Bearer " + token,
	}, nil
}

// Info 获取用户信息
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (l *AdminLogic) AdminInfo(ctx context.Context, req *v1.AdminInfoReq) (res *v1.AdminInfoRes, err error) {
	admin, err := utils.GetAdmin(ctx)
	if err != nil {
		return nil, err
	}
	if admin.Id == 0 {
		return nil, gerror.New("用户不存在")
	}
	if admin.Status == STATUS_DISABLE {
		return nil, gerror.New("用户已禁用")
	}

	data := gconv.Map(admin)
	data["userId"] = data["id"]

	if len(data["real_name"].(string)) > 0 {
		data["userName"] = data["real_name"]
	} else {
		data["userName"] = data["account"]
	}

	data["avatar"] = utils.ImgUrl(data["avatar"].(string))

	delete(data, "password")
	delete(data, "salt")

	return &data, nil
}

// Delete 删除用户
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (l *AdminLogic) AdminDelete(ctx context.Context, req *v1.AdminDeleteReq) (*v1.AdminDeleteRes, error) {
	ids := strings.Split(req.Ids, ",")
	if len(ids) == 0 {
		return nil, gerror.New("用户ID列表不能为空")
	}
	dao.Admins.Ctx(ctx).Where("id in (?)", ids).Delete()
	return &map[string]any{}, nil
}

// List 获取用户列表
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (l *AdminLogic) AdminList(ctx context.Context, req *v1.AdminListReq) (*v1.AdminListRes, error) {
	query := dao.Admins.Ctx(ctx)

	// 构建查询条件
	if req.Account != "" {
		query = query.Where("account like ?", req.Account+"%")
	}
	if req.Phone != "" {
		query = query.Where("phone like ?", req.Phone+"%")
	}
	if req.Email != "" {
		query = query.Where("email like ?", req.Email+"%")
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
	pageRes, err := utils.ToList[ExtAdmins](query, &req.PageReq)
	if err != nil {
		return nil, err
	}
	// 处理图片URL 和角色
	if records, ok := pageRes.Records.([]ExtAdmins); ok {
		if len(records) > 0 {
			l.fillRolesAndAvatar(ctx, records)
		}
	}

	return &v1.AdminListRes{
		PageRes: *pageRes,
	}, nil
}

// fillRolesAndAvatar 填充用户头像和角色
// @param ctx 上下文
// @param admins 用户列表
func (l *AdminLogic) fillRolesAndAvatar(ctx context.Context, admins []ExtAdmins) {
	var adminIds []int
	for i := range admins {
		adminIds = append(adminIds, int(admins[i].Id))
		admins[i].Avatar = utils.ImgUrl(admins[i].Avatar)
	}

	roleAssocs, err := dao.AdminRoleAssocs.Ctx(ctx).
		Where("admin_id in (?)", adminIds).
		All()
	if err != nil {
		return
	}

	var roleIds []int
	adminRoleMap := make(map[int][]int)
	for _, assoc := range roleAssocs {
		adminId := assoc["admin_id"].Int()
		roleId := assoc["role_id"].Int()
		adminRoleMap[adminId] = append(adminRoleMap[adminId], roleId)
		roleIds = append(roleIds, roleId)
	}

	roleMap := make(map[int]map[string]any)
	if len(roleIds) > 0 {
		roles, err := dao.AdminRoles.Ctx(ctx).
			Where("id in (?)", roleIds).
			All()
		if err != nil {
			return
		}
		for _, role := range roles {
			roleMap[role["id"].Int()] = map[string]any{
				"id":   role["id"].Int(),
				"name": role["name"].String(),
			}
		}
	}

	for i := range admins {
		adminId := int(admins[i].Id)
		if roleIds, ok := adminRoleMap[adminId]; ok {
			var roles []map[string]any
			for _, roleId := range roleIds {
				if role, ok := roleMap[roleId]; ok {
					roles = append(roles, role)
				}
			}
			admins[i].Roles = roles
		}
	}
}

// Save 保存用户
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (l *AdminLogic) AdminSave(ctx context.Context, req *v1.AdminSaveReq) (res *v1.AdminSaveRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	data := r.GetMap()
	delete(data, "id")

	account := data["account"].(string)
	if account == "" {
		return nil, gerror.New("账号不能为空")
	}

	if data["password"] != nil && data["password"].(string) != "" {
		data["salt"] = utils.RandomStr(6)
		data["password"] = l.makePassword(data["password"].(string), data["salt"].(string))
	} else {
		delete(data, "password")
		delete(data, "salt")
	}

	data["avatar"] = utils.RemoveDomainPrefix(data["avatar"].(string))

	if req.Id > 0 {
		count, err := dao.Admins.Ctx(ctx).Where("account = ?", account).Where("id <> ?", req.Id).Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, gerror.New("账号已存在")
		}
		_, err = dao.Admins.Ctx(ctx).Where("id = ?", req.Id).Update(data)
	} else {
		count, err := dao.Admins.Ctx(ctx).Where("account = ?", account).Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, gerror.New("账号已存在")
		}
		_, err = dao.Admins.Ctx(ctx).Insert(data)
	}
	if err != nil {
		return nil, err
	}
	return &map[string]any{}, nil
}

// Create 创建用户
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (l *AdminLogic) AdminCreate(ctx context.Context, req *v1.AdminCreateReq) (res *v1.AdminCreateRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	data := r.GetMap()
	delete(data, "id")

	account := data["account"].(string)
	if account == "" {
		return nil, gerror.New("账号不能为空")
	}

	if data["password"] != nil && data["password"].(string) != "" {
		data["salt"] = utils.RandomStr(6)
		data["password"] = l.makePassword(data["password"].(string), data["salt"].(string))
	} else {
		delete(data, "password")
		delete(data, "salt")
	}

	data["avatar"] = utils.RemoveDomainPrefix(data["avatar"].(string))

	if req.Id > 0 {
		count, err := dao.Admins.Ctx(ctx).Where("account = ?", account).Where("id <> ?", req.Id).Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, gerror.New("账号已存在")
		}
		_, err = dao.Admins.Ctx(ctx).Where("id = ?", req.Id).Update(data)
	} else {
		count, err := dao.Admins.Ctx(ctx).Where("account = ?", account).Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, gerror.New("账号已存在")
		}
		_, err = dao.Admins.Ctx(ctx).Insert(data)
	}
	if err != nil {
		return nil, err
	}
	return &map[string]any{}, nil
}

// SetRole 设置用户角色
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (l *AdminLogic) AdminSetRole(ctx context.Context, req *v1.AdminSetRoleReq) (res *v1.AdminSetRoleRes, err error) {
	var admin entity.Admins
	err = dao.Admins.Ctx(ctx).Where("id = ?", req.Id).Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin.Id == 0 {
		return nil, gerror.New("用户不存在")
	}

	roleIds := strings.Split(req.RoleIdIds, ",")

	dao.AdminRoleAssocs.Ctx(ctx).Where("admin_id = ?", req.Id).Delete()

	for _, roleId := range roleIds {
		dao.AdminRoleAssocs.Ctx(ctx).Insert(map[string]any{
			"admin_id": req.Id,
			"role_id":  roleId,
			"created":  utils.Now(),
			"updated":  utils.Now(),
		})
	}

	return &map[string]any{}, nil
}

// createTokenWithTTL 创建token
// @param admin 用户
// @param minutes 过期时间（分钟）
// @return token
// @return error 错误
func (l *AdminLogic) createTokenWithTTL(admin *entity.Admins, minutes int64) (string, error) {
	// 生成token
	uc := &JwtClaims{
		Id:      admin.Id,
		Account: admin.Account,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(minutes) * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(consts.JwtKey))
}

// makePassword 加密密码
// @param password 密码
// @param salt 盐
// @return 加密后的密码
func (l *AdminLogic) makePassword(password string, salt string) string {
	saltMd5, _ := gmd5.Encrypt(salt)
	encrypted, _ := gmd5.Encrypt(saltMd5 + password)
	return encrypted
}
