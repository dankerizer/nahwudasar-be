package fiil

import "nahwudasar-be/pkg/entities"

type InterfaceFiil interface {
	GetAll(page int, filter map[string]string) ([]entities.Fiil, error)
	GetOne(id string) (entities.Fiil, error)
	Count() (int, error)
}
