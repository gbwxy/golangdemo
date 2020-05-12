package abchttp

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"golangdemo/errhandling/filelistingserver/my"
	"net/http"
	"os"
)

type AppHandler func(writer http.ResponseWriter, request *http.Request) error

func ErrWrapper(handler AppHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			r := recover()
			if err, ok := r.(error); ok {
				code := http.StatusInternalServerError
				fmt.Println(err)
				log.Error(err.Error())
				http.Error(writer, http.StatusText(code), code)
			} else {
				str := fmt.Sprintf("Unkown panic type : %v ", r)
				log.Error(str)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Error("Error handling request: %s", err.Error())
			code := http.StatusOK
			if userErr, ok := err.(my.MyError); ok {
				code = http.StatusBadRequest
				http.Error(writer, userErr.Message(), code)
			}

			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}
}
