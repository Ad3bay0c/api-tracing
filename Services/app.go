package Services

import (
	"fmt"
	"github.com/Ad3bay0c/backend-tracing-go/models"
	"github.com/Ad3bay0c/backend-tracing-go/repository"
)

type App interface {
	GetUserInfo() ([]string, error)
	CreateUser(user models.User) error
}

type appService struct {
	repo repository.DB
}

func NewAppService(repo repository.DB) App {
	return appService{
		repo: repo,
	}
}

func (a appService) GetUserInfo() ([]string, error) {
	result := []string{}
	for i := 'a'; i <= 'z'; i++ {
		result = append(result, fmt.Sprintf("Joe %s", string(i)))
	}
	return result, nil
}

func (a appService) CreateUser(user models.User) error {
	return a.repo.CreateUser(user)
}
