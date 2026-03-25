package v1

import "github.com/gogf/gf/v2/frame/g"

type ConfigSetReq struct {
	g.Meta `path:"/config/set" method:"post" summary:"设置配置"`
	Key string `json:"key" dc:"配置键"`
}

type ConfigSetRes = map[string]any

type ConfigGetReq struct {
	g.Meta `path:"/config/get" method:"get" summary:"获取配置"`
	Key string `json:"key" dc:"配置键"`
}

type ConfigGetRes = map[string]any
