package repositories

import (
	"echo-fiber/app/system/models"
	"echo-fiber/pkg/global"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
)

const menus_table = "system_menus"

func GetMenusCount(where *models.WhereMenuParams) (uint64, error) {
	sql := squirrel.Select("count(*)").From(menus_table)
	name := strings.Trim(where.Name, " ")
	remark := strings.Trim(where.Remark, " ")
	status := strings.Trim(where.Status, " ")
	if len(name) > 0 {
		sql = sql.Where(squirrel.Like{"username": fmt.Sprint("%", name, "%")})
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

func GetMenus(where *models.WhereMenuParams) ([]*models.Menu, error) {
	sql := squirrel.Select("id", "name", "parent_id", "path", "permissions", "type", "icon", "status", "sort", "remark", "updated_at").From(menus_table)
	name := strings.Trim(where.Name, " ")
	remark := strings.Trim(where.Remark, " ")
	status := strings.Trim(where.Status, " ")
	if len(name) > 0 {
		sql = sql.Where(squirrel.Like{"username": fmt.Sprint("%", name, "%")})
	}
	if len(remark) > 0 {
		sql = sql.Where(squirrel.Like{"remark": fmt.Sprint("%", remark, "%")})
	}
	if len(status) > 0 {
		sql = sql.Where(squirrel.Eq{"status": status})
	}
	getMenus := sql.OrderBy("created_at").Limit(where.PageSize).Offset((where.Current - 1) * where.PageSize).PlaceholderFormat(squirrel.Question)
	stmt, args, _ := getMenus.ToSql()
	menus := []*models.Menu{}
	err := global.DB.Select(&menus, stmt, args...)
	return menus, err
}

func GetMenu(id string) (*models.Menu, error) {
	sql := squirrel.Select("id", "name", "parent_id", "path", "permissions", "type", "icon", "status", "sort", "remark", "updated_at").From(menus_table)
	stmt, args, _ := sql.Where(squirrel.Eq{"id": id}).PlaceholderFormat(squirrel.Question).ToSql()
	fmt.Println(stmt)
	var menu models.Menu
	err := global.DB.Get(&menu, stmt, args...)
	return &menu, err
}

func CreateMenu(menu *models.Menu) error {
	stmt, args, _ := squirrel.Insert(menus_table).SetMap(squirrel.Eq{
		"id":          menu.ID,
		"name":        menu.Name,
		"parent_id":   menu.ParentId,
		"path":        menu.Path,
		"permissions": menu.Permissions,
		"type":        menu.Type,
		"icon":        menu.Icon,
		"status":      menu.Status,
		"sort":        menu.Sort,
		"remark":      menu.Remark,
		"created_at":  menu.CreatedAt,
		"updated_at":  menu.UpdatedAt,
	}).PlaceholderFormat(squirrel.Question).ToSql()
	_, err := global.DB.Exec(stmt, args...)
	return err
}

func UpdateMenu(updateReq *models.UpdateMenuRequest) error {
	stmt, args, _ := squirrel.Update(menus_table).SetMap(squirrel.Eq{
		"name":        updateReq.Name,
		"parent_id":   updateReq.ParentId,
		"path":        updateReq.Path,
		"permissions": updateReq.Permissions,
		"type":        updateReq.Type,
		"icon":        updateReq.Icon,
		"status":      updateReq.Status,
		"sort":        updateReq.Sort,
		"remark":      updateReq.Remark,
		"updated_at":  updateReq.UpdatedAt,
	}).Where(squirrel.Eq{"id": updateReq.ID}).Limit(1).PlaceholderFormat(squirrel.Question).ToSql()
	_, err := global.DB.Exec(stmt, args...)
	return err
}

func DeleteMenu(id string) error {
	stmt, args, _ := squirrel.Delete(menus_table).Where(squirrel.Eq{"id": id}).PlaceholderFormat(squirrel.Question).ToSql()
	_, err := global.DB.Exec(stmt, args...)
	return err
}

func CheckMenuFields(menu *models.Menu) error {
	id := menu.ID
	name := menu.Name
	path := menu.Path
	// parts := []squirrel.Sqlizer{
	// 	squirrel.Like{}
	// }
	sql := squirrel.Select("id").From(menus_table)
	if len(id) > 0 {
		sql = sql.Where(squirrel.NotEq{"id": id})
		if len(name) > 0 {
			sql = sql.Where(squirrel.Eq{"username": fmt.Sprint("%", name, "%")})
		}
		if len(*path) > 0 {
			sql = sql.Where(squirrel.Eq{"remark": fmt.Sprint("%", path, "%")})
		}
	}
	s, _, err := sql.ToSql()
	fmt.Println(s)
	return err
}
