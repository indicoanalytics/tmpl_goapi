package health

import (
	"time"

	"api.default.indicoinnovation.pt/adapters/database"
	"api.default.indicoinnovation.pt/app/appinstance"
	"api.default.indicoinnovation.pt/entity"
)

type Repository struct {
	database *database.Database[entity.Health]
}

func New() *Repository {
	return &Repository{
		database: database.New[entity.Health](appinstance.Data.DB),
	}
}

func (repo *Repository) Insert(now time.Time) error {
	_, err := repo.database.Exec("INSERT INTO health (sync) VALUES ($1)", now)

	return err
}

func (repo *Repository) GetOne(now time.Time) (*entity.Health, error) {
	return repo.database.QueryOne("SELECT sync FROM health WHERE sync = $1", now)
}
