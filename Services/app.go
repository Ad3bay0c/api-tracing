package Services

type App interface{}

type appService struct{}

func NewAppService() appService {
	return appService{}
}
