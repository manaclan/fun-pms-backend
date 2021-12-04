package hotel

import (
	"manaclan/pms-backend/src/database"

	"github.com/gin-gonic/gin"
)

type HotelRouter struct {
	controllers Controllers
}

func (router *HotelRouter) Init() {
	router.controllers = Controllers{Services: Services{Client: database.GetDB()}}
}

func (hotelRouter HotelRouter) Route(router *gin.RouterGroup) {
	router.POST("/", hotelRouter.controllers.AddHotel)
}
