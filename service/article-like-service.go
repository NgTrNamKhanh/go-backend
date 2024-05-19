package service

import (
	"errors"
	"fmt"

	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/repo"
	"gorm.io/gorm"
)

type ArticleLikeService interface {
	ToggleLikeInArticle(articleLike entity.ArticleLike) (*entity.ArticleLike, error)
	GetLikeCountByArticle(articleID uint) (int64, error)
}

type articleLikeService struct {
}

func NewArticleLikeService() ArticleLikeService {
	return &articleLikeService{}
}

func (s *articleLikeService) ToggleLikeInArticle(articleLike entity.ArticleLike) (*entity.ArticleLike, error) {
	db := repo.NewRepo()

	existingArticleLike, err := db.ArtLik().FindLikeByUserAndArticle(articleLike.UserID, articleLike.ArticleID)
	if err != nil {
		// If the error is not a record not found error, return it
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if existingArticleLike != nil {
		// ArticleLike exists, so delete it
		err = db.ArtLik().DeleteLikeByUserAndArticle(articleLike.UserID, articleLike.ArticleID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		fmt.Println(articleLike)
		// ArticleLike does not exist, so create it
		createdArticleLike, err := db.ArtLik().CreateLike(articleLike)
		if err != nil {
			return nil, err
		}
		return createdArticleLike, nil
	}
}


func (s *articleLikeService) GetLikeCountByArticle(articleID uint) (int64, error) {
	db := repo.NewRepo()
	ArticleLikes, err := db.ArtLik().GetLikeCountByArticle(articleID)
	if err != nil {
		return 0, err
	}
	return ArticleLikes, nil
}
