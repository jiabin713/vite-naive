package services

import (
	"echo-fiber/app/system/models"
	"echo-fiber/app/system/repositories"
	"echo-fiber/pkg/utils"
	"time"
)

func GetStaffs(where *models.WhereStaffParams) ([]*models.Staff, uint64, error) {
	staffs, err := repositories.GetStaffs(where)
	if err != nil {
		return nil, 0, err
	}
	count, err := repositories.GetStaffsCount(where)
	return staffs, count, err
}

func GetStaff(id string) (*models.Staff, error) {
	staff, err := repositories.GetStaff(id)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

func CreateStaff(createReq *models.CreateStaffRequest) error {
	staff := &models.Staff{
		ID:        utils.GenerateID(),
		Username:  createReq.Username,
		Password:  utils.GeneratePassword(createReq.Password),
		Email:     createReq.Email,
		Mobile:    createReq.Mobile,
		Status:    createReq.Status,
		Sort:      createReq.Sort,
		Remark:    createReq.Remark,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	checkStaffFields(staff)
	if err := repositories.CreateStaff(staff); err != nil {
		return err
	}
	return nil
}

func UpdateStaff(updateReq *models.UpdateStaffRequest) error {
	// 检查是否有此id
	if _, err := repositories.GetStaff(updateReq.ID); err != nil {
		return err
	}
	updateReq.UpdatedAt = time.Now().Unix()
	if err := repositories.UpdateStaff(updateReq); err != nil {
		return err
	}
	return nil
}

func DeleteStaff(id string) error {
	if err := repositories.DeleteStaff(id); err != nil {
		return err
	}
	return nil
}

func checkStaffFields(staff *models.Staff) {
	repositories.CheckStaffFields(staff)
}
