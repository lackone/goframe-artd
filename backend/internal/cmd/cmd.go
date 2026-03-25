package cmd

import (
	"backend/internal/consts"
	"backend/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"

	"backend/internal/controller/admin"
)

// initPath 初始化路径
func initPath(s *ghttp.Server) {
	rootDir := gfile.MainPkgPath()

	if rootDir == "" {
		rootDir = gfile.Pwd()
	}

	publicDir := gfile.Join(rootDir, "resource", "public")
	if !gfile.Exists(publicDir) {
		publicDir = "resource/public"
	}

	s.SetServerRoot(publicDir)

	// 添加 /admin 静态路由
	adminDir := gfile.Join(publicDir, consts.AdminDir)
	if gfile.Exists(adminDir) {
		s.AddStaticPath("/admin", adminDir)
	}

	// 添加 uploads 静态路由
	uploadsDir := gfile.Join(publicDir, consts.UploadDir)
	if !gfile.Exists(uploadsDir) {
		gfile.Mkdir(uploadsDir)
	}
	s.AddStaticPath("/"+consts.UploadDir, uploadsDir)
}

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			initPath(s)
			s.Use(MiddlewareCORS, middleware.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.LoginCheck)
					group.Bind(admin.NewV1())
				})
			})
			s.Run()
			return nil
		},
	}
)
