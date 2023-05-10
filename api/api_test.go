package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuliyamun/miniurl/api"
)

func TestAPI_AddUrl(t *testing.T) {
	test := []struct {
		name               string
		handler            api.Handler
		payload            string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "OK",
			handler:            &strHandler{str: "testvalue"},
			payload:            `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`,
			expectedBody:       `{"url":"https://github.com/gourses/miniurl/blob/main/LICENSE", "hash":"testvalue"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Bad request",
			handler:            nil,
			payload:            `invalid json data`,
			expectedBody:       `{"msg":"bad request"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(tc.payload))
			rr := httptest.NewRecorder()

			r := httprouter.New()
			h := &strHandler{str: "testvalue"}
			api.Bind(r, h)
			r.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Result().StatusCode)
			body, err := io.ReadAll(rr.Result().Body)
			require.NoError(t, err)
			assert.JSONEq(t, tc.expectedBody, string(body))
		})

	}
}

type strHandler struct {
	str string
}

func (h *strHandler) AddUrl(url string) (hash string, err error) {
	return h.str, nil
}
