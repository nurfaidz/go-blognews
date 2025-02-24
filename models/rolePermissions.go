package models

type RolePermissions struct {
	GormModel
	RoleID       uint         `gorm:"not null" json:"role_id" form:"role_id" valid:"required~Role ID is required"`
	PermissionID uint         `gorm:"not null" json:"permission_id" form:"permission_id" valid:"required~Permission ID is required"`
	Roles        *Roles       `gorm:"foreignKey:RoleID;references:ID" json:"roles"`
	Permissions  *Permissions `gorm:"foreignKey:PermissionID;references:ID" json:"permissions"`
}
