package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(app *fiber.App) {
	// 压缩
	// app.Use(compress.New(compress.Config{
	// 	Level: compress.LevelBestSpeed, // 1
	// }))
	// 跨域
	app.Use(cors.New())
	// etag
	// app.Use(etag.New())
	// 限流
	// app.Use(limiter.New())
	// logger
	app.Use(logger.New())
	// app.Use(pprof.New())
	// app.Use(recover.New())

}
