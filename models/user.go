package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   string `json:"role_id"`
}

type Role struct {
	ID        uint       `gorm:"primary_key" json:"id_role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	RoleName  string     `json:"role_name" validate:"required"`
	Detail    roleDetail `json:"detail" gorm:"-"`
}

type roleDetail struct {
	Umur string `json:"umur"`
}
