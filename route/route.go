package route

import (
	"dataons-service/middleware"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(Routers *gin.Engine) {
	Routers.Use(requestid.New())
	Routers.NoRoute(middleware.NoRouteHandler())
	Routers.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "GIN GONIC WITH MYSQL TEMPLATE API BY Oviek Shagya")
	})
}
