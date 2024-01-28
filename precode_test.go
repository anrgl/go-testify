package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	URL              string
	TotalCount       int
	ExpectedResponse ExpectedResponse
}

type ExpectedResponse struct {
	StatusCode int
	Body       string
}

func TestMainHandlerWhenStatusOK(t *testing.T) {
	testCases := []TestCase{
		{
			URL: "/cafe?city=moscow&count=4",
			ExpectedResponse: ExpectedResponse{
				StatusCode: http.StatusOK,
				Body:       "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент",
			},
		},
		{
			URL: "/cafe?city=moscow&count=42",
			ExpectedResponse: ExpectedResponse{
				StatusCode: http.StatusOK,
				Body:       "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент",
			},
		},
		{
			URL: "/cafe?city=moscow&count=3",
			ExpectedResponse: ExpectedResponse{
				StatusCode: http.StatusOK,
				Body:       "Мир кофе,Сладкоежка,Кофе и завтраки",
			},
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest("GET", testCase.URL, nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, responseRecorder.Code, testCase.ExpectedResponse.StatusCode)
		body, err := io.ReadAll(responseRecorder.Body)
		require.NoError(t, err)
		assert.Equal(t, string(body), testCase.ExpectedResponse.Body)
	}
}

func TestMainHandlerWhenStatusBadRequest(t *testing.T) {
	testCases := []TestCase{
		{
			URL: "/cafe?city=ufa&count=4",
			ExpectedResponse: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				Body:       WrongCityValue,
			},
		},
		{
			URL: "/cafe?city=moscow",
			ExpectedResponse: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				Body:       CountMissing,
			},
		},
		{
			URL: "/cafe?city=moscow&count=-1",
			ExpectedResponse: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				Body:       WrongCountValue,
			},
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest("GET", testCase.URL, nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)

		require.Equal(t, responseRecorder.Code, testCase.ExpectedResponse.StatusCode)
		body, err := io.ReadAll(responseRecorder.Body)
		require.NoError(t, err)
		assert.Equal(t, string(body), testCase.ExpectedResponse.Body)
	}
}
