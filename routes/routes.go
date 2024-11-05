package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.signup())
	incomingRoutes.POST("/users/signin", controllers.signin())
	incomingRoutes.POST("/admin/add-product", controllers.addProduct()) //admin
	incomingRoutes.GET("/users/products", controllers.getAllProducts())
	incomingRoutes.POST("/users/search", controllers.searchProductByQuery())

}
