package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountEqualTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=4", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	body, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	responseCafeCount := len(strings.Split(string(body), ","))
	assert.Equal(t, responseCafeCount, totalCount)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=42", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	body, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	responseCafeCount := len(strings.Split(string(body), ","))
	assert.Equal(t, responseCafeCount, totalCount)
}

func TestMainHandlerWhenCountLessThanTotal(t *testing.T) {
	totalCount := 3
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=3", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	body, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	responseCafeCount := len(strings.Split(string(body), ","))
	assert.Equal(t, responseCafeCount, totalCount)
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=ufa&count=4", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	body, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	assert.Equal(t, string(body), WrongCityValue)
}

func TestMainHandlerWhenCountMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	body, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	assert.Equal(t, string(body), CountMissing)
}

func TestMainHandlerWhenWrongCountValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=-1", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	body, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	assert.Equal(t, string(body), WrongCountValue)
}
