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
		WHERE sync <> $1
	`, new(entity.Health), "2023-06-09 16:43:56")

	return health, err
}
