package commentLike

import (
	"github.com/NgTrNamKhanh/go-backend/entity"
	"gorm.io/gorm"
)

type Repo interface {
	CreateLike(commentLike entity.CommentLike) (*entity.CommentLike, error)
	GetLikeCountByComment(commentID uint) (int64, error)
	DeleteLikeByUserAndComment(userID uint, commentID uint) error
	FindLikeByUserAndComment(userID uint, commentID uint) (*entity.CommentLike, error)
}

type clRepo struct {
	DB *gorm.DB
}

func NewCommentLikeRepo(db *gorm.DB) Repo {
	return &clRepo{
		DB: db,
	}
}

func (r *clRepo) CreateLike(commentLike entity.CommentLike) (*entity.CommentLike, error) {
	if err := r.DB.Create(&commentLike).Error; err != nil {
		return nil, err
	}
	return &commentLike, nil
}

func (r *clRepo) GetLikeCountByComment(commentID uint) (int64, error) {
	var count int64
	if err := r.DB.Model(&entity.CommentLike{}).Where("comment_id = ?", commentID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}




func (r *clRepo) DeleteLikeByUserAndComment(userID uint, commentID uint) error {
	return r.DB.Where("user_id = ? AND comment_id = ?", userID, commentID).Delete(&entity.CommentLike{}).Error
}
func (r *clRepo) FindLikeByUserAndComment(userID uint, commentID uint) (*entity.CommentLike, error) {
	var commentLike entity.CommentLike
	if err := r.DB.Where("user_id = ? AND comment_id = ?", userID, commentID).First(&commentLike).Error; err != nil {
		return nil, err
	}
	return &commentLike, nil
}