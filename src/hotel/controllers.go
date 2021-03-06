package hotel

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	Services Services
}

/*
 .d8888b.  8888888b.  8888888888        d8888 88888888888 8888888888
d88P  Y88b 888   Y88b 888              d88888     888     888
888    888 888    888 888             d88P888     888     888
888        888   d88P 8888888        d88P 888     888     8888888
888        8888888P"  888           d88P  888     888     888
888    888 888 T88b   888          d88P   888     888     888
Y88b  d88P 888  T88b  888         d8888888888     888     888
 "Y8888P"  888   T88b 8888888888 d88P     888     888     8888888888
*/

func (controllers Controllers) AddHotel(c *gin.Context) {
	fmt.Println("Controller: AddHotel")
	query := c.Request.URL.Query()
	location := Location{
		Address:        query["address"][0],
		Street:         query["street"][0],
		Ward:           query["ward"][0],
		District:       query["district"][0],
		CityOrProvince: query["cityOrProvince"][0],
		Country:        query["country"][0],
	}
	code, message := controllers.Services.AddHotel(query["hotelName"][0], location)
	c.JSON(code, gin.H{
		"status":  "succeed",
		"message": message,
	})
}

/*
8888888b.  8888888888        d8888 8888888b.
888   Y88b 888              d88888 888  "Y88b
888    888 888             d88P888 888    888
888   d88P 8888888        d88P 888 888    888
8888888P"  888           d88P  888 888    888
888 T88b   888          d88P   888 888    888
888  T88b  888         d8888888888 888  .d88P
888   T88b 8888888888 d88P     888 8888888P"
*/

// func (controllers Controllers) Login(c *gin.Context) {
// 	fmt.Println("Controller: Login")
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	isUserLoginSucceed := controllers.Services.ValidateUser(username, password)
// 	fmt.Println(isUserLoginSucceed)
// 	if isUserLoginSucceed {
// 		c.JSON(200, gin.H{
// 			"status":       "succeed",
// 			"message":      "Login succeed!",
// 			"access_token": "jwt-token",
// 		})
// 		return
// 	}
// 	c.JSON(404, gin.H{
// 		"status":  "failed",
// 		"message": "Login failed!",
// 	})
// }

/*
888     888 8888888b.  8888888b.        d8888 88888888888 8888888888
888     888 888   Y88b 888  "Y88b      d88888     888     888
888     888 888    888 888    888     d88P888     888     888
888     888 888   d88P 888    888    d88P 888     888     8888888
888     888 8888888P"  888    888   d88P  888     888     888
888     888 888        888    888  d88P   888     888     888
Y88b. .d88P 888        888  .d88P d8888888888     888     888
 "Y88888P"  888        8888888P" d88P     888     888     8888888888
*/

/*
8888888b.  8888888888 888      8888888888 88888888888 8888888888
888  "Y88b 888        888      888            888     888
888    888 888        888      888            888     888
888    888 8888888    888      8888888        888     8888888
888    888 888        888      888            888     888
888    888 888        888      888            888     888
888  .d88P 888        888      888            888     888
8888888P"  8888888888 88888888 8888888888     888     8888888888
*/
