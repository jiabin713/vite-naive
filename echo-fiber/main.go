package main

import (
	"echo-fiber/pkg/configs"
	"echo-fiber/pkg/global"
	"echo-fiber/pkg/middleware"
	"echo-fiber/pkg/routes"
	"echo-fiber/pkg/utils"
	"echo-fiber/platform/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 读取配置文件
	// configs.ReadEnv()
	// // 连接数据库
	global.DB = database.OpenDBConnection()
	// 迁移数据库
	// migrations.Migrate()
	// fiber 自身配置
	config := configs.FiberConfig()
	// 创建实例
	app := fiber.New(config)
	// 中间件
	middleware.FiberMiddleware(app) // 注册 fiber内置中间件

	// 设置路由
	// routes.SwaggerRoute(app) // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	// if os.Getenv("STAGE_STATUS") == "dev" {
	// 	utils.StartServer(app)
	// } else {
	// 	utils.StartServerWithGracefulShutdown(app)
	// }
	if true {
		fmt.Printf("dev开发中")
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
