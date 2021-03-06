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

func GetStaff(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	staff, err := services.GetStaff(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.OK(staff))
}

func CreateStaff(ctx *fiber.Ctx) error {
	// 新建请求struct 并赋默认值
	createReq := &models.CreateStaffRequest{
		Status: "Default",
		Sort:   1000,
	}
	// 结构体检查 并赋值
	if err := ctx.BodyParser(createReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// 字段检查
	if err := utils.NewValidator().Struct(createReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// 新建实体
	if err := services.CreateStaff(createReq); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// 返回结果
	return ctx.JSON(response.OK(true))
}

func UpdateStaff(ctx *fiber.Ctx) error {
	updateReq := &models.UpdateStaffRequest{}
	if err := ctx.BodyParser(updateReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.NewValidator().Struct(updateReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := services.UpdateStaff(updateReq); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.OK(true))
}

func DeleteStaff(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := services.DeleteStaff(id); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.OK(true))
}
