package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// задание 1. Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerWhenCorrect(t *testing.T) {
	// формирую запрос так, чтобы количество запрашиваемых кафе было меньше равно 4, например, три
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	//получаю ответ от сервера
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//проверка, что код ответа 200
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	// проверка, что тело ответа не пустое
	require.NotNil(t, responseRecorder.Body)

}

// Задание 2. Город, который передаётся в параметре `city`, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку `wrong city value` в теле ответа.
func TestMainHandlerWithWrongCity(t *testing.T) {
	// формирую запрос по городу, который не поддерживается, например, Владивосток
	req := httptest.NewRequest("GET", "/cafe?count=100&city=vladivostok", nil)

	//получаю ответ от сервера
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//проверка, что код ответа 400
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// проверка, что тело ответа возвращает нужное описание ошибки
	errReport := "wrong city value"
	require.Equal(t, errReport, responseRecorder.Body.String())

}

// Задание 3. Если в параметре `count` указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	//создаем запрос и прописываем по условию задания путь /cafe и параметры: количество больше 4 и город Москва
	req := httptest.NewRequest("GET", "/cafe?count=256&city=moscow", nil)

	//получаем ответ от сервера на созданный запрос
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	// перевожу  прошлый  код в систему testify
	//if status := responseRecorder.Code; status != http.StatusOK {t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)}
	// проверка кода ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	// проверка, что тело ответа не пустое
	require.NotNil(t, responseRecorder.Body)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	// проверка, что длина слайса равна максимальному кол-ву кафешек (константе)
	//if len(list) != totalCount {t.Errorf("expected cafe count: %d, got %d", totalCount, len(list))}
	assert.Lenf(t, list, totalCount, "количество кафе не равно %d", totalCount)
	//
}
