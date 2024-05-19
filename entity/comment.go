package entity

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content     string    
	IsAnonymous bool       
	ArticleID   uint       
	ParentID    *uint       
	UserID      uint       
	CommentLikes     []CommentLike  `gorm:"foreignKey:CommentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Replies     []Comment  `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
