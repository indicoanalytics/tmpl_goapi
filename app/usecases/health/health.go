package health

import (
	constants "api.default.indicoinnovation.pt/app/errors"
	"api.default.indicoinnovation.pt/app/repository/health"
	"api.default.indicoinnovation.pt/entity"
)

type Usecase struct {
	repo *health.Repository
}

func New() *Usecase {
	return &Usecase{
		repo: health.New(),
	}
}

func (uc *Usecase) Check() (*entity.Health, error) {
	testDatabase, err := uc.repo.GetHealthCheck()
	if err != nil {
		panic(constants.ErrDatabaseNotConnected)
	}

	return testDatabase, nil
}
