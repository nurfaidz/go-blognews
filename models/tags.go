package models

type Tags struct {
	GormModel
	Name string `gorm:"unique;not null;uniqueIndex" json:"name" form:"name" valid:"required~Your tag name is required"`
	Slug string `gorm:"unique;not null;uniqueIndex" json:"slug" form:"slug" valid:"required~Your tag slug is required"`
}
