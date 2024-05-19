package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/service"
	"github.com/NgTrNamKhanh/go-backend/utils"
	"github.com/gin-gonic/gin"
)

type ArticleLikeController interface {
	ToggleLike(ctx *gin.Context)
	GetLikeCountByArticle(ctx *gin.Context)
}

type articleLikeController struct {
	service service.ArticleLikeService
	validationService service.ValidationService
}

func NewArticleLikeController(service service.ArticleLikeService,validationService service.ValidationService) ArticleLikeController {
	return &articleLikeController{service: service,validationService: validationService}
}

func (c *articleLikeController) ToggleLike(ctx *gin.Context) {
	var like entity.ArticleLike
	if err := ctx.ShouldBindJSON(&like); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := utils.GetUserIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	isUserExist, err := c.validationService.IsUserExist(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user existence"})
        return
    }
    if !isUserExist {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
        return
    }

    isArticleExist, err := c.validationService.IsCommentExist(like.ArticleID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking article existence"})
        return
    }
    if !isArticleExist {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Article not found"})
        return
    }
	like.UserID = userID
	createdLike, err := c.service.ToggleLikeInArticle(like)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add like"})
		return
	}

	ctx.JSON(http.StatusOK, createdLike)
}
func (c *articleLikeController) GetLikeCountByArticle(ctx *gin.Context){
	articleIDParam := ctx.Param("articleID")
    articleID, err := strconv.ParseUint(articleIDParam, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
	isArticleExist, err := c.validationService.IsCommentExist(uint(articleID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking article existence"})
        return
    }
    if !isArticleExist {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Article not found"})
        return
    }
	likeCount, err := c.service.GetLikeCountByArticle(uint (articleID))
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add like"})
		return
	}

	ctx.JSON(http.StatusOK, likeCount)
}

