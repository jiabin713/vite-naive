package models

import (
	"github.com/google/uuid"
)

// entity
type Staff struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"uuid"`
	Username  string    `db:"username" json:"username" validate:"required,min=4,max=64,alphanum"`
	Password  string    `db:"password" json:"-" validate:"required,min=6,max=64"`
	Email     *string   `db:"email" json:"email,omitempty" validate:"required,email"`
	Mobile    *string   `db:"mobile" json:"mobile,omitempty" validate:"required,min=11,max=11"`
	Status    string    `db:"status" json:"status" validate:"required,min=4,max=64"`
	Sort      int32     `db:"sort" json:"sort" validate:"number"`
	Remark    *string   `db:"remark" json:"remark,omitempty"`
	CreatedAt int64     `db:"created_at" json:"-"`
	UpdatedAt int64     `db:"updated_at" json:"updated_at"`
}

// create dto
type CreateStaffRequest struct {
	Username string  `json:"username" validate:"required,min=4,max=64,alphanum"`
	Password string  `json:"password" validate:"required,min=6,max=128"`
	Email    *string `json:"email,omitempty" validate:"required,email"`
	Mobile   *string `db:"mobile" json:"mobile,omitempty" validate:"required,min=11,max=11"`
	Status   string  `db:"status" json:"status" validate:"required,min=4,max=64"`
	Sort     int32   `json:"sort" validate:"number"`
	Remark   *string `json:"remark,omitempty"`
}

// update dto
type UpdateStaffRequest struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Username  string    `db:"username" json:"username" validate:"required,min=4,max=64,alphanum"`
	Email     *string   `db:"email" json:"email,omitempty" validate:"required,email"`
	Mobile    *string   `db:"mobile" json:"mobile,omitempty" validate:"required,min=11,max=11"`
	Status    string    `db:"status" json:"status" validate:"required,min=4,max=64"`
	Sort      int32     `db:"sort" json:"sort" validate:"number"`
	Remark    *string   `db:"remark" json:"remark,omitempty"`
	UpdatedAt int64     `db:"updated_at" json:"updated_at"`
}

// query dto
type WhereStaffParams struct {
	Username string `json:"username"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
	PageSize uint64 `json:"pageSize" validate:"gt=0"`
	Current  uint64 `json:"current" validate:"gt=0"`
}
