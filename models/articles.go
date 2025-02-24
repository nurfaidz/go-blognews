package models

type Articles struct {
	GormModel
	Title      string      `gorm:"not null" json:"title" form:"title" valid:"required~Your article title is required"`
	Slug       string      `gorm:"unique;not null;uniqueIndex" json:"slug" form:"slug" valid:"required~Your article slug is required"`
	Content    string      `gorm:"not null" json:"content" form:"content" valid:"required~Your article content is required"`
	ImageUrl   string      `gorm:"not null" json:"image_url" form:"image_url" valid:"required~Your article image url is required"`
	Status     string      `gorm:"not null" json:"status" form:"status" valid:"required~Your article status is required"`
	CategoryID uint        `gorm:"not null" json:"category_id" form:"category_id" valid:"required~Your article category is required"`
	UserID     uint        `gorm:"not null" json:"user_id" form:"user_id" valid:"required~Your article user is required"`
	Categories *Categories `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	Users      *Users      `gorm:"foreignKey:UserID;references:ID" json:"users"`
}
