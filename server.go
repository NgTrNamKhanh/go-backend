package main

import (
	"github.com/NgTrNamKhanh/go-backend/controller"
	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/initializers"
	"github.com/NgTrNamKhanh/go-backend/middleware"
	"github.com/NgTrNamKhanh/go-backend/repo"
	"github.com/NgTrNamKhanh/go-backend/service"
	"github.com/gin-gonic/gin"
)

var (
	validationService service.ValidationService = service.NewValidationService()

	articleService    service.ArticleService       = service.NewArticleService()
	articleController controller.ArticleController = controller.NewArticleController(articleService,validationService)

	userService    service.UserService       = service.NewUserService()
	userController controller.UserController = controller.NewUserController(userService,validationService)

	commentService   service.CommentService        = service.NewCommentService()
	commentController controller.CommentController = controller.NewCommentController(commentService,validationService)

	articleLikeService   service.ArticleLikeService        = service.NewArticleLikeService()
	articleLikeController controller.ArticleLikeController = controller.NewArticleLikeController(articleLikeService,validationService)

	commentLikeService   service.CommentLikeService        = service.NewCommentLikeService()
	commentLikeController controller.CommentLikeController = controller.NewCommentLikeController(commentLikeService,validationService)
)
func init() {
	initializers.LoadEnvVariables()
}

func main() {
	db := repo.NewRepo()

	db.AutoMigrate(&entity.Article{}, &entity.User{}, &entity.ArticleLike{}, &entity.CommentLike{}, &entity.Comment{})
	
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())

	server.POST("/signup", userController.Signup)
	server.POST("/login", userController.Login)
	server.GET("/validate", middleware.RequireAuth, userController.Validate)

	articlePortal := server.Group("/article", middleware.RequireAuth)
	{
		articlePortal.POST("/create", articleController.CreateArticle)
		articlePortal.GET("/all", articleController.FindAll)
		articlePortal.GET("/user/:userId", articleController.GetArticlesByUser)
	}

	userPortal := server.Group("/user", middleware.RequireAuth)
	{
		userPortal.GET("/all", userController.FindAll)
	}
	commentPortal := server.Group("/comment", middleware.RequireAuth)
	{
		commentPortal.GET("/comments/:articleID", commentController.FindAllParentComment)
		commentPortal.GET("/replies/:commentID", commentController.FindAllReply)
		commentPortal.POST("/", commentController.CreateComment)
		commentPortal.DELETE("/replies", commentController.RemoveComment)
	}
	likePortal := server.Group("/like", middleware.RequireAuth)
	{
		likePortal.GET("/article/:articleID", articleLikeController.GetLikeCountByArticle)
		likePortal.GET("/comment/:commentID", commentLikeController.GetLikeCountByComment)
		likePortal.POST("/article/", articleLikeController.ToggleLike)
		likePortal.POST("/comment/", commentLikeController.ToggleLike)
	}

	server.Run(":8080")

}
