package models

type ArticleTags struct {
	ArticleID uint      `gorm:"not null" json:"article_id" form:"article_id" valid:"required~Article ID is required"`
	TagID     uint      `gorm:"not null" json:"tag_id" form:"tag_id" valid:"required~Tag ID is required"`
	Articles  *Articles `gorm:"foreignKey:ArticleID;references:ID" json:"article"`
	Tags      *Tags     `gorm:"foreignKey:TagID;references:ID" json:"tags"`
}
