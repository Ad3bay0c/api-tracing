package Services

import (
	"context"
	"github.com/Ad3bay0c/backend-tracing-go/application/trace"
	"github.com/Ad3bay0c/backend-tracing-go/models"
	"github.com/Ad3bay0c/backend-tracing-go/repository"
)

type App interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context) ([]models.User, error)
}

type appService struct {
	repo repository.DB
}

func NewAppService(repo repository.DB) App {
	return appService{
		repo: repo,
	}
}

func (a appService) CreateUser(ctx context.Context, user models.User) error {
	ctx, span := trace.NewSpan(ctx, "CreateUser Service Method", nil)
	defer span.End()

	return a.repo.CreateUser(ctx, user)
}

func (a appService) GetUsers(ctx context.Context) ([]models.User, error) {
	ctx, span := trace.NewSpan(ctx, "GetUser Service Method", nil)
	defer span.End()

	return a.repo.GetUsers(ctx)
}
