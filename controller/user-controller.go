package controller

import (
	"net/http"

	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	FindAll(ctx *gin.Context) 
	Signup(ctx *gin.Context) 
	Login(ctx *gin.Context) 
	Validate(ctx *gin.Context)
}

type userController struct {
	service service.UserService
	validationService service.ValidationService
}

func NewUserController(service service.UserService,validationService service.ValidationService) UserController{
	return &userController{
		service: service,
		validationService: validationService,
	}
}

func (c *userController) FindAll(ctx *gin.Context) {
	Users,err := c.service.FindAll()
	if err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Users"})
		return
	}
	ctx.JSON(http.StatusOK, Users)

}
func (c *userController) Signup(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	isEmailExist, err := c.validationService.IsEmailExist(body.Email)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking email existence"})
        return
    }
    if isEmailExist {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email exist"})
        return
    }
	//Hashpassword
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := entity.User{
        Email:    body.Email,
        Password: string(hash),
    }
	createdUser, err := c.service.Add(user)

	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create User"})
        return
    }
    ctx.JSON(http.StatusOK, createdUser)
}

func(c *userController) Login(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	user := entity.User{
		Email:    body.Email,
		Password: body.Password,
	}

	tokenString, err := c.service.Login(user)

	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
        return
    }
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

    ctx.JSON(http.StatusOK, tokenString)


	//Look up requested user
	
}
func(c *userController) Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}