package configs

import (
	"echo-fiber/pkg/exceptions"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {

	return fiber.Config{
		ErrorHandler:          exceptions.ExceptionHandler,
		Prefork:               false,                           // 同一端口多进程
		ServerHeader:          "",                              // 服务报头
		StrictRouting:         false,                           // 严格路由(是否区分 /foo 和 /foo/
		CaseSensitive:         false,                           // 路由区分大小写( /foo 和 /Foo )
		Immutable:             false,                           // 上下文不可变
		UnescapePath:          false,                           // 转换路由编码字符
		ETag:                  false,                           // Etag标头启用
		BodyLimit:             4 * 1024 * 1024,                 // 请求体大小
		Concurrency:           256 * 1024,                      // 并发数
		ReadTimeout:           time.Second * time.Duration(60), // 读取请求允许的时间
		DisableStartupMessage: false,                           // 关闭调试信息
		AppName:               "deer-fiber",                    // APP名称
		EnablePrintRoutes:     false,                           // 打印路由
	}
}
