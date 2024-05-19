package service

import (
	"errors"
	"fmt"

	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/repo"
	"gorm.io/gorm"
)

type CommentLikeService interface {
	ToggleLikeInComment(commentLike entity.CommentLike) (*entity.CommentLike, error)
	GetLikeCountByComment(commentID uint) (int64, error)
}

type commentLikeService struct {
}

func NewCommentLikeService() CommentLikeService {
	return &commentLikeService{}
}

func (s *commentLikeService) ToggleLikeInComment(commentLike entity.CommentLike) (*entity.CommentLike, error) {
	db := repo.NewRepo()

	existingCommentLike, err := db.CmtLik().FindLikeByUserAndComment(commentLike.UserID, commentLike.CommentID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if existingCommentLike != nil {
		err = db.CmtLik().DeleteLikeByUserAndComment(commentLike.UserID, commentLike.CommentID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		fmt.Println(commentLike)
		// commentLike does not exist, so create it
		createdcommentLike, err := db.CmtLik().CreateLike(commentLike)
		if err != nil {
			return nil, err
		}
		return createdcommentLike, nil
	}
}


func (s *commentLikeService) GetLikeCountByComment(commentID uint) (int64, error) {
	db := repo.NewRepo()
	commentLikes, err := db.CmtLik().GetLikeCountByComment(commentID)
	if err != nil {
		return 0, err
	}
	return commentLikes, nil
}


