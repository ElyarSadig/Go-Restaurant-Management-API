package main

import (
	"fmt"
	"log"
	"online_food_market/controller"
	"online_food_market/database"
	"online_food_market/middleware"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadDatabase()
	loadEnv()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	err := database.DB.AutoMigrate(&model.User{}, &model.Customer{}, &model.Order{}, &model.Driver{}, &model.MenuItem{}, &model.DeliveryInfo{}, &model.Restaurant{}, &model.Review{})

	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Successfully Migrated!")
	}
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.GET("/customers", controller.GetAllCustomers)
	protectedRoutes.GET("/customers/:customer_id", controller.GetCustomer)
	protectedRoutes.POST("/customers", controller.CreateCustomer)
	protectedRoutes.PUT("/customers/:customer_id", controller.UpdateCustomer)
	protectedRoutes.DELETE("/customers/:customer_id", controller.DeleteCustomer)

	protectedRoutes.GET("/restaurants", controller.GetAllRestaurants)
	protectedRoutes.POST("/restaurants", controller.CreateRestaurant)
	protectedRoutes.GET("/restaurants/:restaurant_id", controller.GetRestaurant)
	protectedRoutes.PUT("/restaurants/:restaurant_id", controller.UpdateRestaurant)
	protectedRoutes.DELETE("/restaurants/:restaurant_id", controller.DeleteRestaurant)

	protectedRoutes.GET("/menuitems/restaurants/:restaurant_id", controller.GetAllMenuItems)
	protectedRoutes.POST("/menuitems/restaurants/:restaurant_id", controller.CreateMenuItem)
	protectedRoutes.GET("/menuitems/:menuitem_id", controller.GetMenuItem)
	protectedRoutes.PUT("/menuitems/:menuitem_id", controller.UpdateMenuItem)
	protectedRoutes.DELETE("/menuitems/:menuitem_id", controller.DeleteMenuItem)

	protectedRoutes.GET("/orders/customers/:customer_id", controller.GetAllOrders)
	protectedRoutes.POST("/orders/customers/:customer_id", controller.CreateOrder)
	protectedRoutes.GET("/orders/:order_id", controller.GetOrder)
	protectedRoutes.PUT("/orders/:order_id", controller.UpdateOrder)
	protectedRoutes.DELETE("/orders/:order_id", controller.DeleteOrder)

	protectedRoutes.GET("/drivers", controller.GetAllDrivers)
	protectedRoutes.POST("/drivers", controller.CreateDriver)
	protectedRoutes.GET("/drivers/:driver_id", controller.GetDriver)
	protectedRoutes.PUT("/drivers/:driver_id", controller.UpdateDriver)
	protectedRoutes.DELETE("/drivers/:driver_id", controller.DeleteDriver)

	protectedRoutes.GET("/delivery/orders/:order_id", controller.GetDeliveryInfo)
	protectedRoutes.POST("/delivery/orders/:order_id", controller.CreateDeliveryInfo)
	protectedRoutes.PUT("/delivery/:delivery_id", controller.UpdateDeliveryInfo)
	protectedRoutes.DELETE("/delivery/:delivery_id", controller.DeleteDeliveryInfo)

	protectedRoutes.GET("/reviews/restaurants/:restaurant_id", controller.GetAllReviews)
	protectedRoutes.POST("/reviews/restaurants/:restaurant_id", controller.CreateReview)
	protectedRoutes.GET("/reviews/:review_id", controller.GetReview)
	protectedRoutes.PUT("/reviews/:review_id", controller.UpdateReview)
	protectedRoutes.DELETE("/reviews/:review_id", controller.DeleteReview)

	router.Run()

}
