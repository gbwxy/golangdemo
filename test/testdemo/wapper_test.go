package testdemo

import (
	"errors"
	"golangdemo/test/abchttp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func errPanic(w http.ResponseWriter, r *http.Request) error {
	panic(errors.New("123"))
}

var tests = []struct {
	h       abchttp.AppHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
}

//直接调用方法
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		wrapper := abchttp.ErrWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		wrapper(response, request)
		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("except (%d, %s); got (%d, %s)", tt.code, tt.message, response.Code, body)
		}
	}

}

//模拟http server
func TestErrorInServer(t *testing.T) {
	for _, tt := range tests {
		wrapper := abchttp.ErrWrapper(tt.h)

		server := httptest.NewServer(http.HandlerFunc(wrapper))
		response, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.StatusCode != tt.code || body != tt.message {
			t.Errorf("except (%d, %s); got (%d, %s)", tt.code, tt.message, response.StatusCode, body)
		}
	}
}
