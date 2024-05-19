package service

import "github.com/NgTrNamKhanh/go-backend/repo"

type ValidationService interface {
	IsEmailExist(email string) (bool, error)
	IsUserExist(userID uint) (bool, error)
	IsCommentExist(commentID uint) (bool, error)
	IsArticleExist(articleID uint) (bool, error)
}

type validationService struct {
}

func NewValidationService() ValidationService {
	return &validationService{}
}

func (s *validationService) IsUserExist(userID uint) (bool, error) {
    db := repo.NewRepo()

    // Check if the user exists
    exists, err := db.Usr().IsUserExist(userID)
    if err != nil {
        return false, err
    }
    return exists, nil
}

func (s *validationService) IsCommentExist(commentID uint) (bool, error) {
	db := repo.NewRepo()

	exists, err := db.Cmt().IsCommentExist(commentID)
    if err != nil {
        return false, err
    }
    return exists, nil
}
func (s *validationService) IsEmailExist(email string) (bool, error) {
	db := repo.NewRepo()

	exists, err := db.Usr().IsEmailExist(email)
    if err != nil {
        return false, err
    }
    return exists, nil
}
func (s *validationService) IsArticleExist(articleID uint) (bool, error) {
	db := repo.NewRepo()

	exists, err := db.Art().IsArticleExist(articleID)
    if err != nil {
        return false, err
    }
    return exists, nil
}