package main 
import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"strings"
	"mutants/controller"
)

func TestRouter(t *testing.T) {
	cases := []struct {
		name   string
		method string
		url    string
		body   string
	}{
		{
			name:   "POSTMUTANT_OK",
			method: http.MethodPost,
			url:    "/mutant",
			body:   `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`,
		},
		{
			name:   "POSTNOTMUTANT_OK",
			method: http.MethodPost,
			url:    "/mutant",
			body:   `{"dna":["ABCDEF","GHIJKL","MNOPQR","STUVWX","YZABCD","EFGHIJ"]}`,
		},
		/*{
			name:   "GETSTATUS_OK",
			method: http.MethodGet,
			url:    "/stats",
		},*/
	}
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	//_, router := gin.CreateTestContext(w)
	router := newRouter()
	httptest.NewServer(router)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(tc.method, tc.url, strings.NewReader(tc.body))
			r := httptest.NewRecorder()
			router.ServeHTTP(r, req)
			assert.Equal(t, "true", w.Body.String())

			//b := RecordedRespTemplate(t, tc.name, r)

			//AssertWithGolden(t, "mutants", tc.name, b.Bytes())
		})
	}
	
}
func newRouter() (*gin.Engine) {
	router := gin.Default()
	controller.SetControllers(router, nil)
	return router
}