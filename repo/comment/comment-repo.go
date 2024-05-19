package comment

import (
	"github.com/NgTrNamKhanh/go-backend/entity"
	"gorm.io/gorm"
)

// Repo defines the interface for comment repository operations.
type Repo interface {
	FindAllParentComment(articleID uint) ([]entity.Comment, error)
	FindAllReply(parentID uint) ([]entity.Comment, error)
	CreateComment(comment entity.Comment) (*entity.Comment, error)
	DeleteComment(commentID uint) error
	IsCommentExist(commentID uint) (bool, error)
}
// NewCommentRepo creates a new instance of cmtRepo.
func NewCommentRepo(db *gorm.DB) Repo {
	return &cmtRepo{
		DB: db,
	}
}
// cmtRepo is a PostgreSQL implementation of the Repo interface.
type cmtRepo struct {
	DB *gorm.DB
}



// CreateComment creates a new comment in the database.
func (r cmtRepo) CreateComment(comment entity.Comment) (*entity.Comment, error) {
	if err := r.DB.Create(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// DeleteComment deletes a comment from the database by its ID.
func (r cmtRepo) DeleteComment(commentID uint) error {
	if err := r.DB.Delete(&entity.Comment{}, commentID).Error; err != nil {
		return err
	}
	return nil
}

// FindAllParentComment finds all parent comments for a given article ID.
func (r cmtRepo) FindAllParentComment(articleID uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := r.DB.Where("article_id = ? AND parent_id IS NULL", articleID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// FindAllReply finds all replies for a given parent comment ID.
func (r cmtRepo) FindAllReply(parentID uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := r.DB.Where("parent_id = ?", parentID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
func (r cmtRepo) IsCommentExist(commentID uint) (bool, error){
	var comment entity.Comment
	if err := r.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
