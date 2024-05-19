package articleLike

import (
	"github.com/NgTrNamKhanh/go-backend/entity"
	"gorm.io/gorm"
)

type Repo interface {
	CreateLike(articleLike entity.ArticleLike) (*entity.ArticleLike, error)
	GetLikeCountByArticle(articleID uint) (int64, error)
	DeleteLikeByUserAndArticle(userID uint, articleID uint) error
	FindLikeByUserAndArticle(userID uint, articleID uint) (*entity.ArticleLike, error)
}

type alRepo struct {
	DB *gorm.DB
}

func NewArticleLikeRepo(db *gorm.DB) Repo {
	return &alRepo{
		DB: db,
	}
}

func (r *alRepo) CreateLike(articleLike entity.ArticleLike) (*entity.ArticleLike, error) {
	if err := r.DB.Create(&articleLike).Error; err != nil {
		return nil, err
	}
	return &articleLike, nil
}

func (r *alRepo) GetLikeCountByArticle(articleID uint) (int64, error) {
	var count int64
	if err := r.DB.Model(&entity.ArticleLike{}).Where("article_id = ?", articleID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *alRepo) DeleteLikeByUserAndArticle(userID uint, articleID uint) error {
	return r.DB.Where("user_id = ? AND article_id = ?", userID, articleID).Delete(&entity.ArticleLike{}).Error
}
func (r *alRepo) FindLikeByUserAndArticle(userID uint, articleID uint) (*entity.ArticleLike, error) {
	var ArticleLike entity.ArticleLike
	if err := r.DB.Where("user_id = ? AND article_id = ?", userID, articleID).First(&ArticleLike).Error; err != nil {
		return nil, err
	}
	return &ArticleLike, nil
}
