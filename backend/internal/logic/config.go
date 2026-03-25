package logic

import (
	v1 "backend/api/admin/v1"
	"backend/internal/dao"
	"backend/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ConfigLogic struct {
}

func NewConfigLogic() *ConfigLogic {
	return &ConfigLogic{}
}

// ConfigSet 设置配置
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (c *ConfigLogic) ConfigSet(ctx context.Context, req *v1.ConfigSetReq) (*v1.ConfigSetRes, error) {
	r := ghttp.RequestFromCtx(ctx)
	data := r.GetMap()
	key := data["key"].(string)

	delete(data, "key")
	delete(data, "admin_id")

	if len(data) > 0 {
		var insert []map[string]any
		imgFields := garray.NewStrArrayFrom(ImageFields)
		for subKey, value := range data {
			val := fmt.Sprintf("%v", value)
			if imgFields.Contains(subKey) {
				val = utils.RemoveDomainPrefix(val)
			}
			insert = append(insert, map[string]any{
				"key":     key,
				"sub_key": subKey,
				"value":   val,
			})
		}
		dao.Configs.Ctx(ctx).Data(insert).Save()
	}

	return &map[string]any{}, nil
}

// ConfigGet 获取配置
// @param ctx 上下文
// @param req 请求
// @return 响应
// @return error 错误
func (c *ConfigLogic) ConfigGet(ctx context.Context, req *v1.ConfigGetReq) (*v1.ConfigGetRes, error) {
	ret, err := dao.Configs.Ctx(ctx).Fields("value", "sub_key").Where("`key` = ?", req.Key).All()
	if err != nil {
		return nil, err
	}
	result := make(map[string]any)
	imgFields := garray.NewStrArrayFrom(ImageFields)
	for _, item := range ret {
		if imgFields.Contains(item["sub_key"].String()) {
			result[item["sub_key"].String()] = utils.ImgUrl(item["value"].String())
		} else {
			result[item["sub_key"].String()] = item["value"].String()
		}
	}
	return &result, nil
}
