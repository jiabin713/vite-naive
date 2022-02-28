package models

// `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
// `name` varchar(50) DEFAULT NULL COMMENT '菜单名称',
// `parent_id` bigint(20) DEFAULT NULL COMMENT '父菜单ID，一级菜单为0',
// `url` varchar(200) DEFAULT NULL COMMENT '菜单URL,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)',
// `perms` varchar(500) DEFAULT NULL COMMENT '授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)',
// `type` int(11) DEFAULT NULL COMMENT '类型   0：目录   1：菜单   2：按钮',
// `icon` varchar(50) DEFAULT NULL COMMENT '菜单图标',
// `order_num` int(11) DEFAULT NULL COMMENT '排序',
// `create_by` varchar(50) DEFAULT NULL COMMENT '创建人',
// `create_time` datetime DEFAULT NULL COMMENT '创建时间',
// `last_update_by` varchar(50) DEFAULT NULL COMMENT '更新人',
// `last_update_time` datetime DEFAULT NULL COMMENT '更新时间',
// `del_flag` tinyint(4) DEFAULT '0' COMMENT '是否删除  -1：已删除  0：正常',

// MenuLevel     uint                              `json:"-"`
// 	ParentId      string                            `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
// 	Path          string                            `json:"path" gorm:"comment:路由path"`        // 路由path
// 	Name          string                            `json:"name" gorm:"comment:路由name"`        // 路由name
// 	Hidden        bool                              `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
// 	Component     string                            `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
// 	Sort          int                               `json:"sort" gorm:"comment:排序标记"`          // 排序标记
// 	Meta          `json:"meta" gorm:"comment:附加属性"` // 附加属性
// 	SysAuthoritys []SysAuthority                    `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
// 	Children      []SysBaseMenu                     `json:"children" gorm:"-"`
// 	Parameters    []SysBaseMenuParameter            `json:"parameters"`

type Menu struct {
	ID          string  `db:"id" json:"id" validate:""`
	Name        string  `db:"name" json:"name" validate:"required,min=4,max=64"`
	ParentId    string  `db:"parent_id" json:"parent_id" validate:"required"`
	Path        *string `db:"path" json:"path,omitempty" validate:"max=128"`
	Permissions *string `db:"permissions" json:"permissions,omitempty" validate:"max=191"`
	Type        string  `db:"type" json:"type" validate:"alpha,min=4,max=64"`
	Icon        *string `db:"icon" json:"icon,omitempty" validate:""`
	Status      string  `db:"status" json:"status" validate:"required,min=4,max=64"`
	Sort        int32   `db:"sort" json:"sort" validate:"number"`
	Remark      *string `db:"remark" json:"remark,omitempty"`
	CreatedAt   int64   `db:"created_at" json:"-"`
	UpdatedAt   int64   `db:"updated_at" json:"updated_at"`
}

type CreateMenuRequest struct {
	Name        string  `db:"name" json:"name" validate:"required,min=4,max=64"`
	ParentId    string  `db:"parent_id" json:"parent_id" validate:"required"`
	Path        *string `db:"path" json:"path,omitempty" validate:"max=128"`
	Permissions *string `db:"permissions" json:"permissions,omitempty" validate:"max=191"`
	Type        string  `db:"type" json:"type" validate:"alpha,min=4,max=64"`
	Icon        *string `db:"icon" json:"icon,omitempty" validate:""`
	Status      string  `db:"status" json:"status" validate:"required,min=4,max=64"`
	Sort        int32   `db:"sort" json:"sort" validate:"number"`
	Remark      *string `db:"remark" json:"remark,omitempty"`
}

type UpdateMenuRequest struct {
	ID          string  `db:"id" json:"id" validate:"required"`
	Name        string  `db:"name" json:"name" validate:"required,min=4,max=64"`
	ParentId    string  `db:"parent_id" json:"parent_id" validate:"required"`
	Path        *string `db:"path" json:"path,omitempty" validate:"max=128"`
	Permissions *string `db:"permissions" json:"permissions,omitempty" validate:"max=191"`
	Type        string  `db:"type" json:"type" validate:"alpha,min=4,max=64"`
	Icon        *string `db:"icon" json:"icon,omitempty" validate:""`
	Status      string  `db:"status" json:"status" validate:"required,min=4,max=64"`
	Sort        int32   `db:"sort" json:"sort" validate:"number"`
	Remark      *string `db:"remark" json:"remark,omitempty"`
	UpdatedAt   int64   `db:"updated_at" json:"updated_at"`
}

type WhereMenuParams struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
	PageSize uint64 `json:"pageSize" validate:"gt=0"`
	Current  uint64 `json:"current" validate:"gt=0"`
}
