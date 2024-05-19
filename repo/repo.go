package repo

import (
	"github.com/NgTrNamKhanh/go-backend/repo/article"
	"github.com/NgTrNamKhanh/go-backend/repo/articleLike"
	"github.com/NgTrNamKhanh/go-backend/repo/comment"
	"github.com/NgTrNamKhanh/go-backend/repo/commentLike"
	"github.com/NgTrNamKhanh/go-backend/repo/user"
)

type Repo interface {
	AutoMigrate(models ...interface{}) error
	Art() article.Repo
	Usr() user.Repo
	ArtLik() articleLike.Repo
	CmtLik() commentLike.Repo
	Cmt() comment.Repo
}