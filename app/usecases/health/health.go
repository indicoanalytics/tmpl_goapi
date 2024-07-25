package health

import (
	"database/sql"
	"errors"
	"time"

	"api.default.indicoinnovation.pt/app/errs"
	"api.default.indicoinnovation.pt/app/repository/health"
	"api.default.indicoinnovation.pt/entity"
)

var errOutOfSync = errors.New("database is out of sync")

type Usecase struct {
	repo *health.Repository
}

func New() *Usecase {
	return &Usecase{
		repo: health.New(),
	}
}

func (usecase *Usecase) Check() (*entity.Health, error) {
	now := time.Now()

	if err := usecase.repo.Insert(now); err != nil {
		return nil, err
	}

	check, err := usecase.repo.GetOne(now)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &errs.RequestError{Err: errs.ErrHealthNotFound}
		}

		return nil, err
	}

	if check.Sync == nil || check.Sync.IsZero() {
		return nil, &errs.RequestError{Err: errOutOfSync}
	}

	return check, nil
}
