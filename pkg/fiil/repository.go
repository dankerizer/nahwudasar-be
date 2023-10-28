package fiil

import (
	log "github.com/sirupsen/logrus"

	"nahwudasar-be/pkg/entities"

	"gorm.io/gorm"
)

type repositoryFiil struct {
	conn *gorm.DB
}

func (r repositoryFiil) Count() (int, error) {
	var count int64
	err := r.conn.Model(&entities.Fiil{}).Count(&count).Error
	if err != nil {
		log.Errorf("error Count %v", err)
		return -1, err
	}
	return int(count), nil
}

func (r repositoryFiil) GetAll(page int, filter map[string]string) ([]entities.Fiil, error) {
	// var arrFiil []entities.Fiil
	var result []entities.Fiil
	var row int = 20
	tx := r.conn.Model(&entities.Fiil{})

	if page > 0 {
		offset := (page - 1) * row
		tx = tx.Offset(offset).Limit(row)
	}

	if filter != nil {
		tx = tx.Where(filter)
	}

	tx = tx.Order("id asc")

	err := tx.Find(&result).Error

	// err := r.conn.Model(&entities.Fiil{}).Scan(&arrFiil).Error

	if err != nil {
		log.Errorf("Error %v\n", err)
		return nil, err
	}

	return result, nil
}

func (r repositoryFiil) GetOne(id string) (entities.Fiil, error) {
	var fiil entities.Fiil
	err := r.conn.Where("id=?", id).Find(&fiil).Error

	if err != nil {
		log.Printf("Error %v\n", err)
		return entities.Fiil{}, err
	}

	return fiil, err
}

// func newRepo(conn *gorm.DB) InterfaceFiil {
// 	repo := &repository{conn: conn}
// 	return repo
// }

func newRepo(conn *gorm.DB) InterfaceFiil {
	// return &repositoryFiil{conn: conn}
	repo := &repositoryFiil{conn: conn}
	return repo
}
