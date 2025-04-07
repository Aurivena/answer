package answer

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  any `json:"status"`
	Message any `json:"message"`
}

func SendError(c *gin.Context, response string, processStatus ErrorCode) {
	if c == nil {
		return
	}
	status, err := ConvertCodeToStatus(processStatus, StatusCode)
	if err != nil {
		return
	}
	rspError := "Error " + response
	c.JSON(int(processStatus), Response{
		Status:  status,
		Message: rspError,
	})
}

func SendSuccess(c *gin.Context, response string, processStatus ErrorCode) {
	if c == nil {
		return
	}
	status, err := ConvertCodeToStatus(processStatus, StatusCode)
	if err != nil {
		return
	}
	rspSuccess := "Success " + response
	c.JSON(int(processStatus), Response{
		Status:  status,
		Message: rspSuccess,
	})
}

func SendResponseSuccess(c *gin.Context, output any, processStatus ErrorCode) {
	if c == nil {
		return
	}
	status, err := ConvertCodeToStatus(processStatus, StatusCode)
	if err != nil {
		return
	}
	c.JSON(int(processStatus), Response{
		Status:  status,
		Message: output,
	})
}
