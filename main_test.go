package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	//Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) //

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//сервис возвращает код ответа 200
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	//тело ответа не пустое.
	assert.NotEmpty(t, responseRecorder.Body.String())

}

func TestMainHandlerBadCitryRequest(t *testing.T) {
	//Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
	req := httptest.NewRequest("GET", "/cafe?count=2&city=saratov", nil) //

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//сервис возвращает код ответа 400
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	//ошибку wrong city value в теле ответа.
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	//Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) //

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//сервис возвращает код ответа 200
	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	//тело ответа содержит totalCount кафе
	list := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, len(list), totalCount)
}
