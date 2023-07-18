package health

import (
	"api.default.indicoinnovation.pt/adapters/database"
	constants "api.default.indicoinnovation.pt/app/errors"
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

func (repo *Repository) GetHealthCheck() (*entity.Health, error) {
	health, err := repo.db.Query(`
		SELECT *
		FROM health
		WHERE sync <> $1
	`, new(entity.Health), "2023-06-09 16:43:56")

	h, ok := health.(*entity.Health)
	if !ok {
		return nil, constants.ErrAssertDBResponse
	}

	return h, err
}
