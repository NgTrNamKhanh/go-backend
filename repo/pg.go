package repo

import (
	"os"

	"github.com/NgTrNamKhanh/go-backend/repo/article"
	"github.com/NgTrNamKhanh/go-backend/repo/articleLike"
	"github.com/NgTrNamKhanh/go-backend/repo/comment"
	"github.com/NgTrNamKhanh/go-backend/repo/commentLike"
	"github.com/NgTrNamKhanh/go-backend/repo/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgRepo struct {
    DB      *gorm.DB
    Article article.Repo
    User user.Repo
    ArticleLike articleLike.Repo
    CommentLike commentLike.Repo
    Comment comment.Repo
}

// NewRepo creates a new repository
func NewRepo() Repo {
    connectString := os.Getenv("DB_URL")
    db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    return &pgRepo{
        DB:      db,
        Article: article.NewArticleRepo(db),
        User: user.NewUserRepo(db),
        ArticleLike: articleLike.NewArticleLikeRepo(db),
        CommentLike: commentLike.NewCommentLikeRepo(db),
        Comment: comment.NewCommentRepo(db),
    }
}

func (r *pgRepo) AutoMigrate(models ...interface{}) error {
    for idx := range models {
        if err := r.DB.AutoMigrate(models[idx]); err != nil {
            return err
        }
    }
    return nil
}

func (r *pgRepo) Art() article.Repo {
    return r.Article
}

func (r *pgRepo) Usr() user.Repo {
    return r.User
}

func (r *pgRepo) ArtLik() articleLike.Repo {
    return r.ArticleLike
}
func (r *pgRepo) CmtLik() commentLike.Repo {
    return r.CommentLike
}
func (r *pgRepo) Cmt() comment.Repo {
    return r.Comment
}