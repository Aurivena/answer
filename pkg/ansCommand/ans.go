package ans

import (
	"github.com/Aurivena/answer/pkg/ansCode"
	"github.com/Aurivena/answer/pkg/ansError"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  any `json:"status"`
	Message any `json:"message"`
}

func SendError(c *gin.Context, response string, code int) {
	if c == nil {
		return
	}
	status, err := ansCode.ConvertCodeToStatus(code, ansError.StatusCode)
	if err != nil {
		return
	}
	rspError := "Error " + response
	c.JSON(200, Response{
		Status:  status,
		Message: rspError,
	})
}

func SendSuccess(c *gin.Context, response string, code int) {
	if c == nil {
		return
	}
	status, err := ansCode.ConvertCodeToStatus(code, ansError.StatusCode)
	if err != nil {
		return
	}
	rspSuccess := "Success " + response
	c.JSON(200, Response{
		Status:  status,
		Message: rspSuccess,
	})
}

func SendResponseSuccess(c *gin.Context, output any, processStatus int) {
	if c == nil {
		return
	}
	status, err := ansCode.ConvertCodeToStatus(processStatus, ansError.StatusCode)
	if err != nil {
		return
	}
	c.JSON(200, Response{
		Status:  status,
		Message: output,
	})
}
