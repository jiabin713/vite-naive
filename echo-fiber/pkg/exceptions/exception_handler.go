package exceptions

import (
	"github.com/gofiber/fiber/v2"
)

func ExceptionHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	host := ctx.Hostname()
	url := ctx.OriginalURL()

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*Error); ok {
		e.TraceId = url
		e.Host = host
		return ctx.Status(e.ErrorCode).JSON(e)
	}
	if e, ok := err.(*fiber.Error); ok {
		code := e.Code
		return ctx.Status(code).JSON(
			Error{
				Success:      false,
				ErrorCode:    e.Code,
				ErrorMessage: e.Message,
				ShowType:     4,
				TraceId:      url,
				Host:         host,
			},
		)
	}
	// Send custom error page
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}
