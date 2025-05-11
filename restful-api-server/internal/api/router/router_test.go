package router

import (
	"net/http"
	"net/http/httptest"
	"restful-api-server/pkg/logger"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_helloWorld(t *testing.T) {
	var lg = logger.NewLogger()
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
				body:       `{"message":"Hello, World!"}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := NewRouter(lg, nil)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			router.ServeHTTP(w, req)

			if w.Code != tt.want.statusCode {
				t.Errorf("got: %d, want: %d", w.Code, tt.want.statusCode)
			}

			if w.Body.String() != tt.want.body {
				t.Errorf("got: %s, want: %s", w.Body.String(), tt.want.body)
			}
		})
	}
}
