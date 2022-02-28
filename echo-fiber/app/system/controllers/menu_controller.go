package controllers

import (
	"echo-fiber/app/system/models"
	"echo-fiber/app/system/services"
	"echo-fiber/pkg/common/response"
	"echo-fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetMenus(ctx *fiber.Ctx) error {
	where := &models.WhereMenuParams{
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
	menus, count, err := services.GetMenus(where)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.Page(menus, where.Current, where.PageSize, count))
}

func GetMenu(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	menu, err := services.GetMenu(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.OK(menu))
}

func CreateMenu(ctx *fiber.Ctx) error {
	// 新建请求struct 并赋默认值
	createReq := &models.CreateMenuRequest{
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
	if err := services.CreateMenu(createReq); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// 返回结果
	return ctx.JSON(response.OK(true))
}

func UpdateMenu(ctx *fiber.Ctx) error {
	updateReq := &models.UpdateMenuRequest{}
	if err := ctx.BodyParser(updateReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.NewValidator().Struct(updateReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := services.UpdateMenu(updateReq); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.OK(true))
}

func DeleteMenu(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := services.DeleteMenu(id); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.OK(true))
}
