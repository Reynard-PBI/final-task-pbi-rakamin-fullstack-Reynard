package router

import (
	"myproject/controllers"
	"myproject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// User
	router.POST("/users/register", controllers.Register)
	router.POST("/users/login", controllers.Login)
	router.POST("/users/logout", controllers.Logout)
	router.PUT("/users/:userId", middleware.CheckAuth, controllers.UpdateUser)
	router.DELETE("/users/:userId", middleware.CheckAuth, controllers.DeleteUser)

	// Photo
	router.GET("/photos", middleware.CheckAuth, controllers.ShowPhoto)
	router.POST("/photos", middleware.CheckAuth, controllers.CreatePhoto)
	router.PUT("/photos/:photoId", middleware.CheckAuth, controllers.UpdatePhoto)
	router.DELETE("/photos/:photoId", middleware.CheckAuth, controllers.DeletePhoto)

	return router
}