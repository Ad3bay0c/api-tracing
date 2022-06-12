package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Ad3bay0c/backend-tracing-go/Services"
	"github.com/Ad3bay0c/backend-tracing-go/application"
	"github.com/Ad3bay0c/backend-tracing-go/application/trace"
	"github.com/Ad3bay0c/backend-tracing-go/repository"
)

func main() {
	ctx := context.Background()

	prv, err := trace.NewProvider(trace.ProviderConfig{
		Url:            os.Getenv("PROJECT_ID"),
		ServiceName:    "SDS API",
		ServiceVersion: "1.0.0",
		Environment:    "dev",
		Disabled:       false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = prv.Close(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	mux := http.DefaultServeMux

	DB, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	service := Services.NewAppService(DB)
	handler := application.NewHandler(service)

	mux.HandleFunc("/api/v1/create-user", handler.CreateUser)
	mux.HandleFunc("/api/v1/get-user", handler.GetUserInfo)

	log.Println("Server Started")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
