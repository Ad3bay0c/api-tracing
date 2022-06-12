package application

import (
	"github.com/Ad3bay0c/backend-tracing-go/Services"
	"net/http"
)

type handler struct {
	service Services.App
}

func NewHandler(service Services.App) handler {
	return handler{
		service: service,
	}
}

func (h handler) GetInfo(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Working fine and good"))
}
