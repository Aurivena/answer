package test

import (
	"answer/pkg/ansCommand"
	"answer/pkg/ansError"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type test struct {
	Name string
	Age  int
}

var testVal = test{
	Name: "Йога",
	Age:  301,
}

var (
	sendErrorResult           = fmt.Sprintf(`{"status":"%s","message":"Error тест"}`, ansError.StatusCode[ansError.NotFound])
	sendSuccessResult         = fmt.Sprintf(`{"status":"%s","message":"Success тест"}`, ansError.StatusCode[ansError.OK])
	sendResponseSuccessResult = fmt.Sprintf(`{"status": "%s", "message": {"Age": %d, "Name": "%s"}}`, ansError.StatusCode[ansError.OK], testVal.Age, testVal.Name)
)

func TestSendError(t *testing.T) {
	runner.Run(t, "Отправляет json об error", func(t provider.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		ans.SendError(c, "тест", 404)
		assert.JSONEq(t, sendErrorResult, w.Body.String())
	})
}

func TestSendSuccess(t *testing.T) {
	runner.Run(t, "Отправляет json success", func(t provider.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		ans.SendSuccess(c, "тест", 200)
		assert.JSONEq(t, sendSuccessResult, w.Body.String())
	})
}

func TestSendResponseSuccess(t *testing.T) {
	runner.Run(t, "Отправляет структуру Response", func(t provider.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		ans.SendResponseSuccess(c, testVal, 200)
		assert.JSONEq(t, sendResponseSuccessResult, w.Body.String())
	})
}
