package main

import (
	"log"

	"manaclan/pms-backend/src/database"
	"manaclan/pms-backend/src/hotel"
	"manaclan/pms-backend/src/hotels"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	database.Init()
	r := gin.Default()

	hotelRouter := hotel.HotelRouter{}
	hotelRouter.Init()
	hotelRouter.Route(r.Group("/hotel"))
	hotelsRouter := hotels.HotelsRouter{}
	hotelsRouter.Init()
	hotelsRouter.Route(r.Group("/hotels"))
	// usersRouteGroup := r.Group("/users")
	// usersRouter := users.UsersRouter{}
	// usersRouter.Init()
	// usersRouter.Route(usersRouteGroup)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Run(":8525")
}
