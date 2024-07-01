package controller

import (
	"log"
	"net/http"
	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type GoogleController interface {
	StartGoogleAuth(ctx *gin.Context)
	CompleteGoogleAuth(ctx *gin.Context)
}

type googleController struct {
	service service.UserService
}

func NewGoogleController(service service.UserService) GoogleController {
	return &googleController{
		service: service,
	}
}

func (c *googleController) StartGoogleAuth(ctx *gin.Context) {
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (c *googleController) CompleteGoogleAuth(ctx *gin.Context) {
	googleuser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		log.Printf("Google OAuth2 callback failed: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete auth"})
		return
	}

	user := entity.User{
		Email:    googleuser.Email,
		Username: googleuser.Name,
	}
	createdUser, err := c.service.Add(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create User"})
		return
	}

	ctx.JSON(http.StatusCreated, "Successfully signed in ")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Google login successful",
		"user":    createdUser,
	})
}
