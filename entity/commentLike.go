package entity

import (
	"gorm.io/gorm"
)

type CommentLike struct {
	gorm.Model
	UserID    uint
	CommentID uint
}
