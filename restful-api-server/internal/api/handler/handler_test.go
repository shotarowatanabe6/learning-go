package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_helloWorld(t *testing.T) {
	type want struct {
		statusCode int
		body       string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "正常系",
			want: want{
				statusCode: http.StatusOK,
				body:       "Hello, World!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			helloWorld(got, r)

			if got.Code != tt.want.statusCode {
				t.Errorf("status code: %d", got.Code)
			}

			if got.Body.String() != tt.want.body {
				t.Errorf("body: %s", got.Body.String())
			}
		})
	}
}
