package answer

import (
	"github.com/gin-gonic/gin"
)

// Response - Struktur die Antwort
type Response struct {
	Status  any `json:"status"`
	Message any `json:"message"`
}

/*
SendError - dient zu senden eines Fehler code und einer Nachricht.
*/
func SendError(c *gin.Context, response string, processStatus ErrorCode) {
	if c == nil {
		return
	}
	status, err := ConvertCodeToStatus(processStatus)
	if err != nil {
		return
	}
	rspError := "Error " + response
	c.JSON(int(processStatus), Response{
		Status:  status,
		Message: rspError,
	})
}

/*
SendSuccess - dient zu senden eines Erfolg code und einer Nachricht.
*/
func SendSuccess(c *gin.Context, response string, processStatus ErrorCode) {
	if c == nil {
		return
	}
	status, err := ConvertCodeToStatus(processStatus)
	if err != nil {
		return
	}
	rspSuccess := "Success " + response
	c.JSON(int(processStatus), Response{
		Status:  status,
		Message: rspSuccess,
	})
}

/*
SendResponseSuccess - dient zu senden Individuell code.
output - das einer Nachricht oder struktur antwort.
*/
func SendResponseSuccess(c *gin.Context, output any, processStatus ErrorCode) {
	if c == nil {
		return
	}
	status, err := ConvertCodeToStatus(processStatus)
	if err != nil {
		return
	}
	c.JSON(int(processStatus), Response{
		Status:  status,
		Message: output,
	})
}
