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

type CommentLikeController interface {
	ToggleLike(ctx *gin.Context)
	GetLikeCountByComment(ctx *gin.Context)
}

type commentLikeController struct {
	service service.CommentLikeService
	validationService service.ValidationService
}

func NewCommentLikeController(service service.CommentLikeService,validationService service.ValidationService) CommentLikeController {
	return &commentLikeController{service: service, validationService: validationService}
}

func (c *commentLikeController) ToggleLike(ctx *gin.Context) {
    var like entity.CommentLike
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

    isCommentExist, err := c.validationService.IsCommentExist(like.CommentID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking comment existence"})
        return
    }
    if !isCommentExist {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
        return
    }

    like.UserID = userID
    createdLike, err := c.service.ToggleLikeInComment(like)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle like"})
        return
    }

    ctx.JSON(http.StatusOK, createdLike)
}

func (c *commentLikeController) GetLikeCountByComment(ctx *gin.Context){
	commentIDParam := ctx.Param("commentID")
    commentID, err := strconv.ParseUint(commentIDParam, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    isCommentExist, err := c.validationService.IsCommentExist(uint(commentID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking comment existence"})
        return
    }
    if !isCommentExist {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
        return
    }
	likeCount, err := c.service.GetLikeCountByComment(uint (commentID))
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add like"})
		return
	}

	ctx.JSON(http.StatusOK, likeCount)
}
