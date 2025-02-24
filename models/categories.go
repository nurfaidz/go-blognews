package models

type Categories struct {
	GormModel
	Name string `gorm:"unique;not null;uniqueIndex" json:"name" form:"name" valid:"required~Your category name is required"`
}
