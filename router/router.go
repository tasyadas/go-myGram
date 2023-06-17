package router

import (
	"go-myGram/controllers"
	"go-myGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)

	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.POST("/", controllers.CreatePhoto)

		photoRouter.PUT("/:photo_id", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photo_id", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.POST("/", controllers.CreateComment)

		commentRouter.PUT("/:comment_id", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:comment_id", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	sosmedRouter := r.Group("/socmeds")
	{
		sosmedRouter.Use(middlewares.Authentication())
		sosmedRouter.GET("/", controllers.GetAllSosmed)
		sosmedRouter.POST("/", controllers.CreateSocialMedia)

		sosmedRouter.PUT("/:socmed_id", middlewares.SosmedAuthorization(), controllers.UpdateSocialMedia)
		sosmedRouter.DELETE("/:socmed_id", middlewares.SosmedAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
