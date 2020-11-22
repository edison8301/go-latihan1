package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/",HomeHandler)
	r.HandleFunc("/tes",TesHandler)
	http.ListenAndServe(":8080",r)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	nama := request.FormValue("nama")
	writer.WriteHeader(http.StatusOK)
	output := "Hello World: " + nama
	writer.Write([]byte(output))
}

func TesHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Tes"))
}

