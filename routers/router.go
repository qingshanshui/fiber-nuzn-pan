package routers

import (
	v1 "fiber-layout/controllers/v1"
	"fiber-layout/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {
	main := v1.NewDefaultController()
	group := app.Group("/v1")
	group.Post("/list", main.List)        // 获取文件列表
	group.Post("/get", main.Get)          // 获取文件信息
	group.Get("/download", main.Download) // 下载文件
	group.Post("/login", main.Login)      // 登录获取token
	group.Use(middleware.Auth)
	group.Post("/upload", main.Upload) // 上传文件
}
