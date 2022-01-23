package handler

import (
	"crud-challenge/utils"
	"net/http"

	"github.com/golobby/container/v3"
)

func GetHandler(handler Handler) func(http.ResponseWriter,
	*http.Request)  {

	return func(writer http.ResponseWriter, request *http.Request) {
		err := container.Fill(handler)
		if err != nil {
			utils.OutputInternalServerError(writer,err)
		} else {
			handler.Handle(writer, request)
		}
	}
}
