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

type ArticleController interface {
	FindAll(ctx *gin.Context) 
	CreateArticle(ctx *gin.Context) 
    GetArticlesByUser(ctx *gin.Context)
}

type articleController struct {
	service service.ArticleService
    validationService service.ValidationService
}

func NewArticleController(service service.ArticleService,validationService service.ValidationService) ArticleController{
	return &articleController{
		service: service,
        validationService: validationService,
	}
}

func (c *articleController) GetArticlesByUser(ctx *gin.Context) {
    userIDParam := ctx.Param("userId")
    userID, err := strconv.ParseUint(userIDParam, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    isUserExist, err := c.validationService.IsUserExist(uint(userID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user existence"})
        return
    }
    if !isUserExist {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
        return
    }
    articles, err := c.service.GetArticlesByUserID(uint(userID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
        return
    }
    ctx.JSON(http.StatusOK, articles)
}

func (c *articleController) FindAll(ctx *gin.Context) {
	articles,err := c.service.FindAll()
	if err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles"})
		return
	}
	ctx.JSON(http.StatusOK, articles)

}

// Save handles POST requests to save a new article
func (c *articleController) CreateArticle(ctx *gin.Context) {
    var article entity.Article

    // Extract the user ID from the JWT token
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
    // Bind JSON request to the article struct
    if err := ctx.ShouldBindJSON(&article); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Add the article with the user ID
	
    createdArticle, err := c.service.Add(article, userID)
    fmt.Println(err)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
        return
    }

    ctx.JSON(http.StatusOK, createdArticle)
}
