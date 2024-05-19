package article

import (
	"github.com/NgTrNamKhanh/go-backend/entity"
	"gorm.io/gorm"
)

type Repo interface {
	CreateArticle(p entity.Article) (*entity.Article, error)
	UpdateArticle(ID string, newA entity.Article) (*entity.Article, error)
	GetArticleByID(ID string) (*entity.Article, error)
	GetArticleList() ([]entity.Article, error)
	DeleteArticle(ID string) error
	GetArticlesByUserID(userID uint) ([]entity.Article, error)
	IsArticleExist(articleID uint) (bool, error)
}

func NewArticleRepo(db *gorm.DB) Repo {
	return &aRepo{
		DB: db,
	}
}

type aRepo struct {
	DB *gorm.DB
}

func (r aRepo) UpdateArticle(ID string, newa entity.Article) (*entity.Article, error) {
	var a entity.Article
	if err := r.DB.First(&a, ID).Error; err != nil {
		return nil, err
	}

	a.Title = newa.Title
	a.Description = newa.Description
	a.View = newa.View
	a.IsAnonymous = newa.IsAnonymous
	a.Status = newa.Status

	if err := r.DB.Save(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

// CreateArticle implements Repo.
func (r aRepo) CreateArticle(a entity.Article) (*entity.Article, error) {
	return &a, r.DB.Create(&a).Error
}

// DeleteArticle implements Repo.
func (r aRepo) DeleteArticle(ID string) error {
	p := entity.Article{}
	return r.DB.Table("articles").Where("id = ?", ID).Delete(&p).Error
}

// GetArticleByID implements Repo.
func (r aRepo) GetArticleByID(ID string) (*entity.Article, error) {
	p := entity.Article{}
	return &p, r.DB.Table("articles").Where("id = ?", ID).First(&p).Error
}
func (r aRepo) GetArticlesByUserID(userID uint) ([]entity.Article, error) {
    var articles []entity.Article
    if err := r.DB.Where("user_id = ?", userID).Find(&articles).Error; err != nil {
        return nil, err
    }
    return articles, nil
}
func (r aRepo) GetArticleList() ([]entity.Article, error) {
	rs := []entity.Article{}
	return rs, r.DB.Table("articles").Find(&rs).Error
}

func (r aRepo) IsArticleExist(articleID uint) (bool, error){
	var article entity.Article
	if err := r.DB.Where("id = ?", articleID).First(&article).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
