package main

import (
	"github.com/Ad3bay0c/backend-tracing-go/Services"
	"github.com/Ad3bay0c/backend-tracing-go/application"
	"log"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	service := Services.NewAppService()
	handler := application.NewHandler(service)

	mux.HandleFunc("/", handler.GetInfo)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
