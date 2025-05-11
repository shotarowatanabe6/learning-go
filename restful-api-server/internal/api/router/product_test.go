package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"restful-api-server/internal/models"
	"restful-api-server/internal/service"
	"restful-api-server/pkg/logger"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_GetProducts(t *testing.T) {
	var lg = logger.NewLogger()
	type want struct {
		statusCode int
		body       models.Products
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "正常系",
			want: want{
				statusCode: http.StatusOK,
				body: models.Products{
					{ID: 1, Name: "Product 1", Price: 100},
					{ID: 2, Name: "Product 2", Price: 200},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			svc := service.NewService(nil)
			router := NewRouter(lg, svc)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/products", nil)
			router.ServeHTTP(w, req)

			if w.Code != tt.want.statusCode {
				t.Errorf("got: %d, want: %d", w.Code, tt.want.statusCode)
			}

			var got models.Products
			err := json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Errorf("failed to unmarshal response body: %v, body: %s", err, w.Body.String())
			}

			if !reflect.DeepEqual(got, tt.want.body) {
				t.Errorf("got: %+v, want: %+v", got, tt.want.body)
			}
		})
	}
}
