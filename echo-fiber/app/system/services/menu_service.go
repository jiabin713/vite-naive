package services

import (
	"echo-fiber/app/system/models"
	"echo-fiber/app/system/repositories"
	"echo-fiber/pkg/utils"
	"time"
)

func GetMenus(where *models.WhereMenuParams) ([]*models.Menu, uint64, error) {
	menus, err := repositories.GetMenus(where)
	if err != nil {
		return nil, 0, err
	}
	count, err := repositories.GetMenusCount(where)
	return menus, count, err
}

func GetMenu(id string) (*models.Menu, error) {
	menu, err := repositories.GetMenu(id)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func CreateMenu(createReq *models.CreateMenuRequest) error {
	menu := &models.Menu{
		ID:          utils.GenerateID(),
		Name:        createReq.Name,
		ParentId:    createReq.ParentId,
		Path:        createReq.Path,
		Permissions: createReq.Permissions,
		Type:        createReq.Type,
		Icon:        createReq.Icon,
		Status:      createReq.Status,
		Sort:        createReq.Sort,
		Remark:      createReq.Remark,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	checkMenuFields(menu)
	if err := repositories.CreateMenu(menu); err != nil {
		return err
	}
	return nil
}

func UpdateMenu(updateReq *models.UpdateMenuRequest) error {
	// 检查是否有此id
	if _, err := repositories.GetMenu(updateReq.ID); err != nil {
		return err
	}
	updateReq.UpdatedAt = time.Now().Unix()
	if err := repositories.UpdateMenu(updateReq); err != nil {
		return err
	}
	return nil
}

func DeleteMenu(id string) error {
	if err := repositories.DeleteMenu(id); err != nil {
		return err
	}
	return nil
}

func checkMenuFields(menu *models.Menu) {
	repositories.CheckMenuFields(menu)
}
