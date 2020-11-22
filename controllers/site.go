package controllers

import (
	"net/http"
)

func TesHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Tes"))
}