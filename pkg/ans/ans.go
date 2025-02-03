package ans

import (
	"answer/pkg/ansCode"
	"answer/pkg/ansError"
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
	status, err := ansError.ConvertCodeToStatus(code, ansCode.StatusCode)
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
	status, err := ansError.ConvertCodeToStatus(code, ansCode.StatusCode)
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
	status, err := ansError.ConvertCodeToStatus(processStatus, ansCode.StatusCode)
	if err != nil {
		return
	}
	c.JSON(200, Response{
		Status:  status,
		Message: output,
	})
}
