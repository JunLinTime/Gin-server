package initialize

import (
	"server/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("")
	{
		systemRouter.InitUserRouter(PublicGroup)
	}
	return Router
}
