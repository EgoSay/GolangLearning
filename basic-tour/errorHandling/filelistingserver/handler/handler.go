/**
 * @Author codeAC
 * @Time: 2018/8/9 15:08
 * @Description
 */
package handler

import (
	"log"
	"net/http"
	"os"
)

type UserError interface {
	error
	Message() string
}

type FileHandler func(writer http.ResponseWriter, request *http.Request) error

func ErrorWrapper(handler FileHandler) func(writer http.
	ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Print(1, "Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Print(2,
				"Error occurred "+
					"handling request: %s",
				err.Error())
			//user error
			if useErr, ok := err.(UserError); ok {
				http.Error(writer,
					useErr.Message(),
					http.StatusBadRequest)
				return
			}
			//system error
			code := http.StatusOK
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
