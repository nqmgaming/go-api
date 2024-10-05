package routes

import (
	"github.com/gin-gonic/gin"
	controllers "go-api/controllers"
	middleware "go-api/middleware"
)

func SetupRouter() *gin.Engine {
	request := gin.Default()

	// Routes for authentication
	request.POST("/register", controllers.Register)
	request.POST("/login", controllers.Login)

	// Routes for books (protected by JWT)
	bookRoutes := request.Group("/books")
	bookRoutes.Use(middleware.AuthMiddleware())
	{
		bookRoutes.GET("/", controllers.GetBooks)
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}

	return request
}
