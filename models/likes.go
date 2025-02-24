package models

type Likes struct {
	GormModel
	UserID    uint      `json:"user_id"`
	ArticleID uint      `json:"article_id"`
	Articles  *Articles `gorm:"foreignKey:ArticleID;references:ID" json:"article"`
	Users     *Users    `gorm:"foreignKey:UserID;references:ID" json:"users"`
}
