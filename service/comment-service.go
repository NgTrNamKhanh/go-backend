package service

import (
	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/repo"
)

type CommentService interface {
	FindAllParentComment(articleID uint) ([]entity.Comment, error)
	FindAllReply(parentID uint) ([]entity.Comment, error)
	CreateComment(comment entity.Comment,  userID uint) (*entity.Comment, error)
	RemoveComment(commentID uint) error
}

type commentService struct {
}

func NewCommentService() CommentService {
	return &commentService{}
}

func (s *commentService) CreateComment(comment entity.Comment,  userID uint) (*entity.Comment, error) {
	db := repo.NewRepo()

	newComment := entity.Comment{
        Content:    comment.Content,
        IsAnonymous: comment.IsAnonymous,
		ArticleID: comment.ArticleID,
		ParentID: comment.ParentID,
		UserID: userID,
    }
	// Like does not exist, so create it
	createdComment, err := db.Cmt().CreateComment(newComment)
	if err != nil {
		return nil, err
	}
	return createdComment, nil
}
func (s *commentService)FindAllParentComment(articleID uint) ([]entity.Comment, error){
	db := repo.NewRepo()
	comments, err := db.Cmt().FindAllParentComment(articleID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
	
func (s *commentService)FindAllReply(parentID uint) ([]entity.Comment, error){
	db := repo.NewRepo()
	comments, err := db.Cmt().FindAllReply(parentID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *commentService) RemoveComment(commentID uint) error {
	db := repo.NewRepo()
	return db.Cmt().DeleteComment(commentID)
}

