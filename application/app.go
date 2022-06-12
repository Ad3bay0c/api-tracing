package application

import (
	"encoding/json"
	"fmt"
	"github.com/Ad3bay0c/backend-tracing-go/application/trace"
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
	ctx := req.Context()

	ctx, span := trace.NewSpan(ctx, "CreateUser DB Method", nil)
	defer span.End()

	user := models.User{
		ID:    "1",
		Name:  "Joe",
		Email: "joe@example.com",
	}
	err := h.service.CreateUser(ctx, user)
	if err != nil {
		trace.AddSpanError(span, err)
		trace.FailSpan(span, fmt.Sprintf("internal server error: %v", err.Error()))

		json.NewEncoder(w).Encode("an error occurred: " + err.Error())
		return
	}
	json.NewEncoder(w).Encode("successful")
}

func (h handler) GetUserInfo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx, span := trace.NewSpan(ctx, "Get User Handler", nil)
	defer span.End()

	users, err := h.service.GetUsers(ctx)
	if err != nil {
		trace.AddSpanError(span, err)
		trace.FailSpan(span, fmt.Sprintf("internal server error: %v", err.Error()))

		json.NewEncoder(w).Encode("an error occurred: " + err.Error() + span.SpanContext().TraceID().String())
		return
	}
	json.NewEncoder(w).Encode(users)
}
