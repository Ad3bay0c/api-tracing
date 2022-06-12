package Services

import (
	"github.com/Ad3bay0c/backend-tracing-go/models"
	"github.com/Ad3bay0c/backend-tracing-go/repository"
)

type App interface {
	CreateUser(user models.User) error
	GetUsers() ([]models.User, error)
}

type appService struct {
	repo repository.DB
}

func NewAppService(repo repository.DB) App {
	return appService{
		repo: repo,
	}
}

func (a appService) CreateUser(user models.User) error {
	return a.repo.CreateUser(user)
}

func (a appService) GetUsers() ([]models.User, error) {
	return a.repo.GetUsers()
}
