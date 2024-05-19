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

type CommentController interface {
	FindAllParentComment(ctx *gin.Context) 
	FindAllReply(ctx *gin.Context) 
	CreateComment(ctx *gin.Context) 
	CreateReply(ctx *gin.Context) 
	RemoveComment(ctx *gin.Context)
}

type commentController struct {
	service service.CommentService
	validationService service.ValidationService
}

func NewCommentController(service service.CommentService,validationService service.ValidationService) CommentController{
	return &commentController{
		service: service,
		validationService: validationService,
	}
}

func (c *commentController) FindAllParentComment(ctx *gin.Context) {
	articleID, err := strconv.ParseUint(ctx.Param("articleID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid like ID"})
		return
	}
	isArticleExist, err := c.validationService.IsArticleExist(uint(articleID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking article existence"})
        return
    }
    if !isArticleExist {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Article not found"})
        return
    }
	comments,err := c.service.FindAllParentComment(uint(articleID))
	if err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}
	ctx.JSON(http.StatusOK, comments)

}
func (c *commentController) FindAllReply(ctx *gin.Context) {
	commentID, err := strconv.ParseUint(ctx.Param("commentID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid like ID"})
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
	comments,err := c.service.FindAllReply(uint(commentID));
	if err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}
	ctx.JSON(http.StatusOK, comments)

}
// Save handles POST requests to save a new comment
func (c *commentController) CreateComment(ctx *gin.Context) {
    var comment entity.Comment

    // Extract the user ID from the JWT token
    userID, err := utils.GetUserIDFromToken(ctx)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Bind JSON request to the comment struct
    if err := ctx.ShouldBindJSON(&comment); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	fmt.Println(comment.ArticleID)
    // Add the comment with the user ID
	
    createdComment, err := c.service.CreateComment(comment, userID)
    fmt.Println(err)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
        return
    }

    ctx.JSON(http.StatusOK, createdComment)
}

func (c *commentController) CreateReply(ctx *gin.Context) {
    var reply entity.Comment

    // Extract the user ID from the JWT token
    userID, err := utils.GetUserIDFromToken(ctx)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Bind JSON request to the comment struct
    if err := ctx.ShouldBindJSON(&reply); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Add the comment with the user ID
	
    createdComment, err := c.service.CreateComment(reply, userID)
    fmt.Println(err)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
        return
    }

    ctx.JSON(http.StatusOK, createdComment)
}

func (c *commentController) RemoveComment(ctx *gin.Context) {
	commentID, err := strconv.ParseUint(ctx.Param("commentID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid like ID"})
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
	if err := c.service.RemoveComment(uint(commentID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove like"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Like removed"})
}