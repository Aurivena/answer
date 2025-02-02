package answer

import (
	"answer/pkg/err"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Status  any `json:"status"`
	Message any `json:"message"`
}

func SendError(c *gin.Context, message string, status any) {
	msgError := "Error" + message
	c.JSON(200, Message{
		Status:  status,
		Message: msgError,
	})
}

func SendSuccess(c *gin.Context, message string, status any) {
	msgSuccess := "Success" + message
	c.JSON(200, Message{
		Status:  status,
		Message: msgSuccess,
	})
}

func SendResponseSuccess(c *gin.Context, output any, processStatus int) {
	if c == nil {
		return
	}
	status := err.ConvertCodeToStatus(processStatus)
	c.JSON(200, Message{
		Status:  status,
		Message: output,
	})
}
