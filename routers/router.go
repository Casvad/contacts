package routers

import (
	"contacts/controllers"
	"contacts/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router interface {
	CreateRouter() *gin.Engine
}

type router struct {
	userController        controllers.UserController
	userContactController controllers.UserContactController
}

func ProvideRouter(userController controllers.UserController,
	userContactController controllers.UserContactController,
) Router {

	return &router{userController, userContactController}
}

func (r *router) CreateRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middlewares.ErrorHandler())

	contactRouter := ginRouter.Group("/api/contacts")
	{

		contactRouter.GET("/health", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"status": "OK"})
		})

		userRouter := contactRouter.Group("/users")
		{
			userRouter.POST("/login", r.userController.Login)
			userRouter.POST("/register", r.userController.Register)
		}

		userContactRouter := contactRouter.Group("/contacts")
		userContactRouter.Use(middlewares.AuthMiddleware())
		{
			userContactRouter.POST("/", r.userContactController.Create)
			userContactRouter.PUT("/:id", r.userContactController.Update)
			userContactRouter.DELETE("/:id", r.userContactController.Delete)
			userContactRouter.GET("", r.userContactController.GetAllForUser)
			userContactRouter.GET("/:id", r.userContactController.GetById)
		}

	}

	return ginRouter
}
