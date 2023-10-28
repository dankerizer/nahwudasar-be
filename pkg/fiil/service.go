package fiil

import (
	"nahwudasar-be/pkg/entities"

	"gorm.io/gorm"
)

type service struct {
	repoFiil InterfaceFiil
}

// Count implements InterfaceFiil.
func (s *service) Count() (int, error) {
	return s.repoFiil.Count()
}

func (s service) GetAll(page int, filter map[string]string) ([]entities.Fiil, error) {
	return s.repoFiil.GetAll(page, filter)
}

func (s service) GetOne(id string) (entities.Fiil, error) {
	return s.repoFiil.GetOne(id)
}

func NewService(conn *gorm.DB) InterfaceFiil {
	return &service{repoFiil: newRepo(conn)}
}
