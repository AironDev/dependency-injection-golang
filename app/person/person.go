package app

import (
	"github.com/airondev/dependency-injection-golang/domain/repositories"
)

type PersonApp struct {
	repositories.PersonRepository
}

func (r repositories.PersonRepository) FindAll() {

}
