package models

type Roles struct {
	GormModel
	Name string `gorm:"unique;not null;uniqueIndex" json:"name" form:"name" valid:"required~Your role name is required"`
}
