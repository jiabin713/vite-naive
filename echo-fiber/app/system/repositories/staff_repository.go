package repositories

import (
	"echo-fiber/app/system/models"
	"echo-fiber/pkg/global"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
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
	fmt.Println(where.Status)
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
