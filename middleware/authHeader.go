package middleware

import (
	"dataons-service/pkg"
	"dataons-service/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "The processing function of the request route was not found"})
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-key, x-signature, x-timestamp, x-access, x-key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := response.Gin{C: c}
		key := c.GetHeader("x-key")
		if key != pkg.KEYHEADER {
			appG.Response(http.StatusUnauthorized, "Auth Failed", "", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := response.Gin{C: c}
		key := c.GetHeader("x-access")
		redisCon := pkg.InitializeRedis()
		var data pkg.GenerateAccess
		if err := redisCon.GetKey(key, &data); err != nil {
			appG.Response(http.StatusUnauthorized, "Access Header Failed", "", nil)
			c.Abort()
			return
		}
		c.Next()

	}
}

func AuthBasic() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := response.Gin{C: c}
		user, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && user == pkg.USERNAME && password == pkg.PASSWORD {
			//fmt.Println("NEXT BASIC AUTH")
			c.Next()
		} else {
			appG.Response(http.StatusBadRequest, "", "Basic Auth Failed", nil)
			c.Abort()
			return
		}

	}
}
