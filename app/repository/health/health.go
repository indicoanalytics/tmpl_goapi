package health

import (
	"time"

	"api.default.indicoinnovation.pt/adapters/database"
	constantserrors "api.default.indicoinnovation.pt/app/errors"
	"api.default.indicoinnovation.pt/entity"
)

type Repository struct {
	db *database.Database
}

func New() *Repository {
	return &Repository{
		db: database.New(),
	}
}

func (repo *Repository) Insert(now time.Time) error {
	return repo.db.Exec("INSERT INTO health (sync) VALUES ($1)", now)
}

func (repo *Repository) GetOne(now time.Time) (*entity.Health, error) {
	health, err := repo.db.QueryOne("SELECT * FROM health WHERE sync = $1", &entity.Health{}, now)

	h, ok := health.(*entity.Health)
	if !ok {
		return nil, constantserrors.ErrAssertDBResponse
	}

	return h, err
}
