package controllers

import (
	"echo-fiber/app/system/models"
	"echo-fiber/app/system/services"
	"echo-fiber/pkg/common/response"
	"echo-fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetStaffs(ctx *fiber.Ctx) error {
	where := &models.WhereStaffParams{
		PageSize: 10,
		Current:  1,
	}
	// 解析
	if err := ctx.QueryParser(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// 字段检查
	if err := utils.NewValidator().Struct(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	staffs, count, err := services.GetStaffs(where)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.Page(staffs, where.Current, where.PageSize, count))
}
