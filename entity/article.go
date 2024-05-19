package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string 
	Description string 
	View        int   
	IsAnonymous bool   
	Status      int   
	UserID      uint
	Comments       []Comment `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ArticleLikes       []ArticleLike `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
