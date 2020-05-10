package filelisting

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func FileListingHandle(writer http.ResponseWriter, request *http.Request) {

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	i, err := writer.Write(all)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}

const preFix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func FileListingHandleThrows(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, preFix) != 0 {
		//return errors.New("path must start with : " + preFix)
		return userError("path must start with : " + preFix)
	}

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		log.Warn("Path is : %s", path)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	i, err := writer.Write(all)
	if err != nil {
		return err
	}
	fmt.Println(i)
	return nil
}
