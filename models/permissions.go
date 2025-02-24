package models

type Permissions struct {
	GormModel
	Name string `gorm:"unique;not null;uniqueIndex" json:"name" form:"name" valid:"required~Your permission name is required"`
}
