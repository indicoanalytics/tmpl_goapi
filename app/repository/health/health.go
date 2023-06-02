package health

import (
	"api.default.indicoinnovation.pt/adapters/database"
	"api.default.indicoinnovation.pt/entity"
)

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (repo *Repository) GetHealthCheck() (*entity.Health, error) {
	health, err := database.Query(`
		SELECT *
		FROM health
	`, new(entity.Health))

	return health.(*entity.Health), err
}
