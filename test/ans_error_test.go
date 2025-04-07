package test

import (
	"github.com/Aurivena/answer"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/stretchr/testify/assert"
	"testing"
)

var test_statusCode = map[answer.ErrorCode]string{
	1: "Test1",
}

func f() answer.ErrorCode {
	return answer.OK
}

func TestAppendCode(t *testing.T) {
	processStatus := f()
	answer.SendResponseSuccess(nil, "", processStatus)
	runner.Run(t, "Проверяет, что добавляется новый код в test_statusCode.", func(t provider.T) {
		_ = answer.AppendCode(2, "Test2", test_statusCode)
		if ok, _ := answer.ConvertCodeToStatus(2, test_statusCode); ok == "Test2" {
			assert.True(t, true, "Новый код добавлен в map.")
		} else {
			assert.False(t, false, "Ошибка. Метод не добавляет код в map.")
		}
	})
}

func TestFailAppendCode(t *testing.T) {
	runner.Run(t, "Проверяет, что новый код не будет добавлен в test_statusCode и вызовет ошибку.", func(t provider.T) {
		err := answer.AppendCode(1, "Test2", test_statusCode)
		if err != nil {
			assert.True(t, true, "Код обработал событие при котором текущий код уже существует.")
		} else {
			assert.True(t, false, "Ошибка. Код некорректно обрабатывает событие при котором такой код уже существует.")
		}

	})
}

func TestDeleteCode(t *testing.T) {
	runner.Run(t, "Проверяет, что удаляет код в test_statusCode.", func(t provider.T) {
		_ = answer.DeleteCode(2, test_statusCode)
		if ok, _ := answer.ConvertCodeToStatus(2, test_statusCode); ok != "Test1" {
			assert.True(t, true, "Код удален из map.")
		} else {
			assert.False(t, false, "Ошибка. Метод не удаляет код из map.")
		}
	})
}

func TestFailDeleteCode(t *testing.T) {
	runner.Run(t, "Проверяет, что функция не удалит код в test_statusCode и вызовет ошибку.", func(t provider.T) {
		err := answer.DeleteCode(666, test_statusCode)
		if err != nil {
			assert.True(t, true, "Код корректно обрабатывает событие при котором код не будет удален, если его нету в map")
		} else {
			assert.False(t, false, "Ошибка. Код не обрабатывает событие при котором код не будет удален, если его нету в map")
		}
	})
}

func TestConvertCodeToStatus(t *testing.T) {
	runner.Run(t, "Проверяет, что конвертирует указанный код в статус.", func(t provider.T) {
		if ok, _ := answer.ConvertCodeToStatus(1, test_statusCode); ok == "Test1" {
			assert.True(t, true, "Функция правильно конвертирует указанный код в статус.")
		} else {
			assert.False(t, false, "Ошибка. Метод не конвертирует указанный код в статус.")
		}
	})
}

func TestFailConvertCodeToStatus(t *testing.T) {
	runner.Run(t, "Проверяет, что функция не конвертирует указанный код в статус и вызовет ошибку.", func(t provider.T) {
		_, err := answer.ConvertCodeToStatus(666, test_statusCode)
		if err != nil {
			assert.True(t, true, "Конвертация не произошла тк такого кода не существует")
		} else {
			assert.False(t, false, "Ошибка. Неправильно обрабатывает событие при котором в map нету такого значения")
		}
	})
}
