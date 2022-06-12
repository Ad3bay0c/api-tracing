package application

import (
	"encoding/json"
	"net/http"

	"github.com/Ad3bay0c/backend-tracing-go/Services"
	"github.com/Ad3bay0c/backend-tracing-go/models"
)

type handler struct {
	service Services.App
}

func NewHandler(service Services.App) handler {
	return handler{
		service: service,
	}
}

func (h handler) CreateUser(w http.ResponseWriter, req *http.Request) {
	user := models.User{
		ID:    "1",
		Name:  "Joe",
		Email: "joe@example.com",
	}
	err := h.service.CreateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode("an error occurred: " + err.Error())
	}
	json.NewEncoder(w).Encode("successful")
}
