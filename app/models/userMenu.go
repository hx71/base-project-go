package models

import (
	"time"

	"gorm.io/gorm"
)

type UserMenu struct {
	ID        uint           `gorm:"primary_key:auto_increment" json:"id"`
	RoleID    uint           `json:"role_id" binding:"required"`
	MenuID    uint           `json:"menu_id" binding:"required"`
	Read      string         `gorm:"not null; default:'1'" json:"read"`
	Create    string         `gorm:"not null; default:'1'" json:"create"`
	Update    string         `gorm:"not null; default:'1'" json:"update"`
	Delete    string         `gorm:"not null; default:'1'" json:"delete"`
	Report    string         `gorm:"not null; default:'1'" json:"report"`
	Role      Role           `json:"role" binding:"required"`
	Menu      Menu           `json:"menu" binding:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
