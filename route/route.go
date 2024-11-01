package route

import (
	"dataons-service/controllers"
	"dataons-service/middleware"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(Routers *gin.Engine) {
	Routers.Use(requestid.New(), middleware.CORSMiddleware(), middleware.AuthHeader())
	Routers.NoRoute(middleware.NoRouteHandler())
	Routers.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "GIN GONIC WITH MYSQL TEMPLATE API BY Oviek Shagya")
	})

	master := Routers.Group("/master")
	master.Use(middleware.AuthBasic())
	{
		master.GET("/:idCompany/:idDepart/:idDivision", controllers.UserController.MasterDataCompany)
		master.GET("/inher", controllers.UserController.MasterCompanyInheritance)
		master.POST("/company", controllers.UserController.CreateUpdateCompany)
		master.DELETE("/company", controllers.UserController.DeleteCompany)
	}

}
