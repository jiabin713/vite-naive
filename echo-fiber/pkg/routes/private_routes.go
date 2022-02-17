package routes

import (
	"echo-fiber/app/system/controllers"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/")
	SystemRoutes(route)
}

func SystemRoutes(route fiber.Router) {
	system := route.Group("/system")
	StaffRoutes(system)
	// DictionaryRoutes(system)
	// DictionaryItemRoutes(system)
}

func StaffRoutes(route fiber.Router) {
	staffs := route.Group("staffs")
	staffs.Get("/", controllers.GetStaffs)
	// staffs.Post("/", controllers.CreateStaff)
	// staffs.Put("/", controllers.UpdateStaff)
	// staffs.Get("/:id", controllers.GetStaff)
	// staffs.Delete("/:id", controllers.DeleteStaff)
}

// func DictionaryRoutes(route fiber.Router) {
// 	dictionaries := route.Group("dictionaries")
// 	dictionaries.Get("/", controllers.GetDictionaries)
// 	dictionaries.Post("/", controllers.CreateDictionary)
// 	dictionaries.Put("/", controllers.UpdateDictionary)
// 	dictionaries.Get("/:id", controllers.GetDictionary)
// 	dictionaries.Delete("/:id", controllers.DeleteDictionary)
// }

// func DictionaryItemRoutes(route fiber.Router) {
// 	dictionaries := route.Group("dictionary/items")
// 	dictionaries.Get("/", controllers.GetDictionaryItems)
// 	dictionaries.Post("/", controllers.CreateDictionaryItem)
// 	dictionaries.Put("/", controllers.UpdateDictionaryItem)
// 	dictionaries.Get("/:id", controllers.GetDictionaryItem)
// 	dictionaries.Delete("/:id", controllers.DeleteDictionaryItem)
// }
