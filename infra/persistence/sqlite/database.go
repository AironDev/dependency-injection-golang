package sqlite

import (
	"database/sql"

	"github.com/airondev/dependency-injection-golang/domain/repositories"
	"github.com/airondev/dependency-injection-golang/infra/config"
)

type Repositories struct {
	Person   repositories.PersonRepository
	database *sql.DB
}

func InitDatabase(c *config.Config) (*Repositories, error) {
	db, _ := sql.Open("sqlite3", c.DatabasePath)
	return &Repositories{
		Person:   NewPersonRepository(db),
		database: db,
	}, nil
}
