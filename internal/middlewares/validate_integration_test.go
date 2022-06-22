package middlewares

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/one-byte-data/go-api-sample/cmd/tests"
	"gorm.io/driver/postgres"
)

func TestIntegrationValidateHeader(t *testing.T) {
	if m := flag.Lookup("test.run").Value.String(); m == "" || !regexp.MustCompile(m).MatchString(t.Name()) {
		t.Skip("skipping as execution was not requested explicitly using go test -run")
	}
	
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.Use(ValidateHeader())

	health := router.Group("/health")
	{
		health.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
	}

	type args struct {
		method   string
		endpoint string
		header   map[string]string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "Should check health",
			args: args{
				method:   "GET",
				endpoint: "/health",
				header:   map[string]string{},
			},
			wantCode: http.StatusOK,
		},
		{
			name: "Should not be able check health",
			args: args{
				method:   "GET",
				endpoint: "/health",
				header: map[string]string{
					"X-Not-Valid": "true",
				},
			},
			wantCode: http.StatusForbidden,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.args.method, tt.args.endpoint, nil)
		for key, value := range tt.args.header {
			req.Header.Add(key, value)
		}
		router.ServeHTTP(w, req)

		if tt.wantCode != w.Code {
			t.Errorf("ValidateHeader() error = %v, wantCode %v", w.Code, tt.wantCode)
			return
		}
	}
}
