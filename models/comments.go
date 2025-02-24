package models

type Comments struct {
	GormModel
	Content   string    `gorm:"not null" json:"content" form:"content" valid:"required~Your comment content is required"`
	ArticleID uint      `json:"article_id" gorm:"not null" form:"article_id" valid:"required~Article ID is required"`
	UserID    uint      `gorm:"not null" json:"user_id" form:"user_id" valid:"required~User ID is required"`
	Articles  *Articles `gorm:"foreignKey:ArticleID;references:ID" json:"article"`
	Users     *Users    `gorm:"foreignKey:UserID;references:ID" json:"users"`
}
