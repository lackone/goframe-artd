package utils

import (
	"backend/api/common"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

// MakeTree 将扁平数组转换为树形结构
// @param items 扁平数组
// @param id 主键字段名
// @param pid 父键字段名
// @param son 子项字段名
// @return 树形结构数组
func MakeTree(items []map[string]any, id string, pid string, son string) []map[string]any {
	if len(items) == 0 {
		return nil
	}

	// 建立 id -> index 的映射
	mp := make(map[string]int)
	for i := range items {
		mp[fmt.Sprintf("%v", items[i][id])] = i
	}

	// 初始化 children
	for i := range items {
		items[i][son] = []map[string]any{}
	}

	// 构建树关系
	var roots []map[string]any
	for i := range items {
		pidVal := items[i][pid]
		pidKey := fmt.Sprintf("%v", pidVal)

		// 根节点：pid 为 0、空、或者找不到父节点
		if pidKey == "0" || pidKey == "" {
			roots = append(roots, items[i])
		} else if parentIdx, ok := mp[pidKey]; ok {
			// 找到父节点，加入父节点的 children
			items[parentIdx][son] = append(items[parentIdx][son].([]map[string]any), items[i])
		} else {
			// 找不到父节点，作为根节点
			roots = append(roots, items[i])
		}
	}

	return roots
}

// ToTimestamp 时间字符串转换为时间戳
// @param timeStr 时间字符串
// @return 时间戳
func ToTimestamp(timeStr string) int64 {
	return gtime.NewFromStr(timeStr).Timestamp()
}

// ToList 分页查询
// @param query 查询对象
// @param req 分页请求
// @return 分页结果
func ToList[T any](query *gdb.Model, req *common.PageReq) (*common.PageRes, error) {
	// 统计总记录数
	total, _ := query.Count()

	// 分页
	if req.Page > 0 && req.Size > 0 {
		query = query.Page(req.Page, req.Size)
	}

	// 查询
	var list []T
	err := query.Scan(&list)
	if err != nil {
		return nil, err
	}

	return &common.PageRes{
		Current: req.Page,
		Records: list,
		Total:   int(total),
		Size:    req.Size,
	}, nil
}

// RandomStr 随机字符串
// @param n 随机字符串长度
// @return 随机字符串
func RandomStr(n int) string {
	return grand.Str("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", n)
}

// Now 当前时间戳
// @return 当前时间戳
func Now() int64 {
	return time.Now().Unix()
}

// GetAdmin 获取当前登录管理员
// @param ctx 上下文
// @return 管理员
func GetAdmin(ctx context.Context) (*entity.Admins, error) {
	var admin entity.Admins
	adminId := ghttp.RequestFromCtx(ctx).GetCtxVar("admin_id").Uint()
	err := dao.Admins.Ctx(ctx).Where("id = ?", adminId).Scan(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// Asset 资产路径
// @param path 路径
// @return 资产路径
func Asset(path string) string {
	if path == "" {
		return ""
	}
	return strings.TrimRight(consts.Domain, "/") + "/" + strings.TrimLeft(path, "/")
}

// ImgUrl 图片路径
// @param path 路径
// @return 图片路径
func ImgUrl(path string) string {
	if path == "" {
		return ""
	}
	re := regexp.MustCompile(`^https?://`)
	url := re.ReplaceAllString(Asset(path), "")
	return "//" + url
}

// RemoveDomainPrefix 移除域名前缀
// @param urlStr URL字符串
// @return 移除域名前缀后的URL字符串
func RemoveDomainPrefix(urlStr string) string {
	parsed, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}
	path := parsed.Path
	if strings.HasPrefix(path, "/" + consts.UploadDir + "/") {
		return path
	}
	pos := strings.Index(strings.ToLower(urlStr), "/" + consts.UploadDir + "/")
	if pos != -1 {
		return urlStr[pos:]
	}
	return urlStr
}