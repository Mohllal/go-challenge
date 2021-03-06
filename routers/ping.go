package routers

import (
	"github.com/Mohllal/go-challenge/controllers"
	"github.com/gin-gonic/gin"
)

var PingRouter = func(baseUrl string, route *gin.Engine) {
	home := route.Group(baseUrl)
	{
		home.GET("/ping", controllers.GetPong)
	}
}
