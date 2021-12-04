package hotels

import (
	"manaclan/pms-backend/src/database"

	"github.com/gin-gonic/gin"
)

type HotelsRouter struct {
	controllers Controllers
}

func (router *HotelsRouter) Init() {
	router.controllers = Controllers{Services: Services{Client: database.GetDB()}}
}

func (hotelsRouter HotelsRouter) Route(router *gin.RouterGroup) {
	// router.POST("/", hotelRouter.controllers.AddHotel)
	router.GET("/", hotelsRouter.controllers.SearchByKeyword)

}
