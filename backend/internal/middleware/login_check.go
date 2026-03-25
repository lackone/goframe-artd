package middleware

import (
	"backend/internal/consts"
	"backend/internal/logic"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

var noAuthPaths = []string{
	"/api/admin/login",
	"/api/common/auth_list",
	"/api/common/role_list",
	"/api/common/upload",
}

// LoginCheck 登录检查
func LoginCheck(r *ghttp.Request) {
	for _, path := range noAuthPaths {
		if r.URL.Path == path {
			r.Middleware.Next()
			return
		}
	}
	var tokenString = r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := &logic.JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(consts.JwtKey), nil
	})
	if err != nil || !token.Valid {
		r.Response.WriteStatus(401)
		r.Response.WriteJsonExit(map[string]any{
			"code": 401,
			"msg":  "token无效",
			"data": nil,
		})
		return
	}
	if claims.Id == 0 {
		r.Response.WriteStatus(401)
		r.Response.WriteJsonExit(map[string]any{
			"code": 401,
			"msg":  "token无效",
			"data": nil,
		})
		return
	}
	r.SetCtxVar("admin_id", claims.Id)

	r.Middleware.Next()
}
