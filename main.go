package main

import (
	"log"
	"net/http"

	"github.com/Ad3bay0c/backend-tracing-go/Services"
	"github.com/Ad3bay0c/backend-tracing-go/application"
	"github.com/Ad3bay0c/backend-tracing-go/repository"
)

func main() {
	mux := http.DefaultServeMux

	DB, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	service := Services.NewAppService(DB)
	handler := application.NewHandler(service)

	mux.HandleFunc("/api/v1/create-user", handler.CreateUser)

	log.Println("Server Started")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
