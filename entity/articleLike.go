package entity

import (
	"gorm.io/gorm"
)

type ArticleLike struct {
	gorm.Model
	UserID    uint
	ArticleID uint
}
