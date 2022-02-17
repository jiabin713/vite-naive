package repositories

import (
	"echo-fiber/app/system/models"
	"echo-fiber/pkg/global"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const staffs_table = "system_staffs"

func GetStaffsCount(where *models.WhereStaffParams) (uint64, error) {
	sql := squirrel.Select("count(*)").From(staffs_table)
	username := strings.Trim(where.Username, " ")
	remark := strings.Trim(where.Remark, " ")
	status := strings.Trim(where.Status, " ")
	if len(username) > 0 {
		sql = sql.Where(squirrel.Like{"username": fmt.Sprint("%", username, "%")})
	}
	if len(remark) > 0 {
		sql = sql.Where(squirrel.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	if len(status) > 0 {
		sql = sql.Where(squirrel.Eq{"status": status})
	}
	getCount := sql.PlaceholderFormat(squirrel.Question)
	stmt, args, _ := getCount.ToSql()

	var count uint64
	err := global.DB.QueryRowx(stmt, args...).Scan(&count)
	return count, err
}

func GetStaffs(where *models.WhereStaffParams) ([]*models.Staff, error) {
	sql := squirrel.Select("id", "username", "email", "mobile", "status", "sort", "remark", "updated_at").From(staffs_table)
	username := strings.Trim(where.Username, " ")
	remark := strings.Trim(where.Remark, " ")
	status := strings.Trim(where.Status, " ")
	if len(username) > 0 {
		sql = sql.Where(squirrel.Like{"username": fmt.Sprint("%", username, "%")})
	}
	if len(remark) > 0 {
		sql = sql.Where(squirrel.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	if len(status) > 0 {
		sql = sql.Where(squirrel.Eq{"status": status})
	}
	getStaffs := sql.OrderBy("created_at").Limit(where.PageSize).Offset((where.Current - 1) * where.PageSize).PlaceholderFormat(squirrel.Question)
	stmt, args, _ := getStaffs.ToSql()
	staffs := []*models.Staff{}
	err := global.DB.Select(&staffs, stmt, args...)
	return staffs, err
}

func GetStaff(id uuid.UUID) (*models.Staff, error) {
	sql := squirrel.Select("id", "username", "email", "mobile", "status", "sort", "remark", "updated_at").From(staffs_table)
	stmt, args, _ := sql.Where(squirrel.Eq{"id": id}).PlaceholderFormat(squirrel.Question).ToSql()
	var staff models.Staff
	err := global.DB.Get(&staff, stmt, args...)
	return &staff, err
}

func CreateStaff(staff *models.Staff) error {
	stmt, args, _ := squirrel.Insert(staffs_table).SetMap(squirrel.Eq{
		"id":         staff.ID,
		"username":   staff.Username,
		"password":   staff.Password,
		"email":      staff.Email,
		"mobile":     staff.Mobile,
		"status":     staff.Status,
		"sort":       staff.Sort,
		"remark":     staff.Remark,
		"created_at": staff.CreatedAt,
		"updated_at": staff.UpdatedAt,
	}).PlaceholderFormat(squirrel.Question).ToSql()
	_, err := global.DB.Exec(stmt, args...)
	return err
}

func UpdateStaff(updateReq *models.UpdateStaffRequest) error {
	stmt, args, _ := squirrel.Update(staffs_table).SetMap(squirrel.Eq{
		"username":   updateReq.Username,
		"email":      updateReq.Email,
		"mobile":     updateReq.Mobile,
		"status":     updateReq.Status,
		"sort":       updateReq.Sort,
		"remark":     updateReq.Remark,
		"updated_at": updateReq.UpdatedAt,
	}).Where(squirrel.Eq{"id": updateReq.ID}).Limit(1).PlaceholderFormat(squirrel.Question).ToSql()
	_, err := global.DB.Exec(stmt, args...)
	return err
}

func DeleteStaff(id uuid.UUID) error {
	stmt, args, _ := squirrel.Delete(staffs_table).Where(squirrel.Eq{"id": id}).PlaceholderFormat(squirrel.Question).ToSql()
	_, err := global.DB.Exec(stmt, args...)
	return err
}

func CheckStaffFields(staff *models.Staff) error {
	id := staff.ID
	username := staff.Username
	email := staff.Email
	mobile := staff.Mobile
	// parts := []squirrel.Sqlizer{
	// 	squirrel.Like{}
	// }
	sql := squirrel.Select("id").From(staffs_table)
	if len(id) > 0 {
		sql = sql.Where(squirrel.NotEq{"id": id})
		if len(username) > 0 {
			sql = sql.Where(squirrel.Eq{"username": fmt.Sprint("%", username, "%")})
		}
		if len(*email) > 0 {
			sql = sql.Where(squirrel.Eq{"remark": fmt.Sprint("%", email, "%")})
		}
		if len(*mobile) > 0 {
			sql = sql.Where(squirrel.Eq{"remark": fmt.Sprint("%", email, "%")})
		}
	}
	s, _, err := sql.ToSql()
	fmt.Println(s)
	return err
}
