package main

import (
	"buynownow/controllers"
	"buynownow/database"
	"buynownow/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.authenticaton())

	router.POST("/add-to-cart", app.AddToCart())
	router.POST("/remove-item", app.RemoveItem())
	router.POST("/checkout", app.BuyFromCart())
	router.POST("/instant-buy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
