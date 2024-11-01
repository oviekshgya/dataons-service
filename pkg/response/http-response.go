package response

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Gin struct {
	C *gin.Context
}

type ResponseSukses struct {
	Code         int         `json:"code"`
	Message      string      `json:"message,omitempty"`
	TimeStamps   time.Time   `json:"timeStamps"`
	Errormessage string      `json:"errorMessage,omitempty"`
	Data         interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, message, errorMessage string, data interface{}) {
	g.C.JSON(httpCode, ResponseSukses{
		TimeStamps:   time.Now(),
		Code:         httpCode,
		Data:         data,
		Errormessage: errorMessage,
		Message:      message,
	})
	return

}
