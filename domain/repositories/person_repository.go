package repositories

import (
	"github.com/airondev/dependency-injection-golang/domain/entities"
)

type PersonRepository interface {
	FindAll() []*entities.Person
}
