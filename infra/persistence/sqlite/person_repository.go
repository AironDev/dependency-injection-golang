package sqlite

import (
	"database/sql"

	"github.com/airondev/dependency-injection-golang/domain/entities"
	"github.com/airondev/dependency-injection-golang/domain/repositories"
)

type PersonRepository struct {
	repositories.PersonRepository
	database *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}

func (repository *PersonRepository) FindAll() []*entities.Person {
	rows, _ := repository.database.Query(
		`SELECT id, name, age FROM people;`,
	)
	defer rows.Close()

	people := []*entities.Person{}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		rows.Scan(&id, &name, &age)

		people = append(people, &entities.Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}

	return people
}
