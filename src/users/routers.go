package users

import (
	"manaclan/pms-backend/src/database"

	"github.com/gin-gonic/gin"
)

type UsersRouter struct {
	Services    Services
	Controllers Controllers
}

func (router UsersRouter) Init() {
	router.Services = Services{Client: database.DB}
	router.Controllers = Controllers{Services: router.Services}
}

func (ur UsersRouter) Route(router *gin.RouterGroup) {
	router.POST("/login", ur.Controllers.Login)
	router.POST("/register", ur.Controllers.Register)
}
