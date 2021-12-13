package user

import (
	"manaclan/pms-backend/src/database"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Services    Services
	Controllers Controllers
}

func (router *UserRouter) Init() {
	router.Services = Services{Client: database.DB}
	router.Controllers = Controllers{Services: router.Services}
}

func (ur *UserRouter) Route(router *gin.RouterGroup) {
	router.POST("/login", ur.Controllers.Login)
	router.POST("/register", ur.Controllers.Register)
}
