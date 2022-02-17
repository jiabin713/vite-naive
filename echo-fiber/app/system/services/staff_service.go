package services

import (
	"echo-fiber/app/system/models"
	"echo-fiber/app/system/repositories"
)

func GetStaffs(where *models.WhereStaffParams) ([]*models.Staff, uint64, error) {
	staffs, err := repositories.GetStaffs(where)
	if err != nil {
		return nil, 0, err
	}
	count, err := repositories.GetStaffsCount(where)
	return staffs, count, err
}
