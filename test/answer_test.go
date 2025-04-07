package test

import (
	"fmt"
	"github.com/Aurivena/answer"
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
	sendErrorResult           = fmt.Sprintf(`{"status":"%s","message":"Error тест"}`, answer.StatusCode[answer.NotFound])
	sendSuccessResult         = fmt.Sprintf(`{"status":"%s","message":"Success тест"}`, answer.StatusCode[answer.OK])
	sendResponseSuccessResult = fmt.Sprintf(`{"status": "%s", "message": {"Age": %d, "Name": "%s"}}`, answer.StatusCode[answer.OK], testVal.Age, testVal.Name)
)

func TestSendError(t *testing.T) {
	runner.Run(t, "Отправляет json об error", func(t provider.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		answer.SendError(c, "тест", answer.NotFound)
		assert.JSONEq(t, sendErrorResult, w.Body.String())
	})
}

func TestSendSuccess(t *testing.T) {
	runner.Run(t, "Отправляет json success", func(t provider.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		answer.SendSuccess(c, "тест", answer.OK)
		assert.JSONEq(t, sendSuccessResult, w.Body.String())
	})
}

func TestSendResponseSuccess(t *testing.T) {
	runner.Run(t, "Отправляет структуру Response", func(t provider.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		answer.SendResponseSuccess(c, testVal, answer.OK)
		assert.JSONEq(t, sendResponseSuccessResult, w.Body.String())
	})
}
