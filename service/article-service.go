package service

import (
	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/repo"
)

type ArticleService interface {
	Add(article entity.Article,  userID uint) (*entity.Article, error)
	FindAll() ([]entity.Article, error)
	GetArticlesByUserID(userID uint) ([]entity.Article, error)
}

type articleService struct {
}

func NewArticleService() ArticleService {
	return &articleService{}
}

func (s *articleService) Add(article entity.Article,  userID uint) (*entity.Article, error) {
	db := repo.NewRepo()
	newArticle := entity.Article{
        Title:    article.Title,
        Description: article.Description,
		View: 0,
		Status: 1,
		IsAnonymous: article.IsAnonymous,
		UserID: userID,
    }
	createdArticle, err := db.Art().CreateArticle(newArticle)
	if err != nil {
		return nil, err
	}
	return createdArticle, nil
}

func (s *articleService) FindAll() ([]entity.Article, error) {
	db := repo.NewRepo()
	articles, err := db.Art().GetArticleList()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *articleService) GetArticlesByUserID(userID uint) ([]entity.Article, error) {
	db := repo.NewRepo()
    articles, err := db.Art().GetArticlesByUserID(userID)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
