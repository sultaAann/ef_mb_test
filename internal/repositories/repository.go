package repositories

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll()
	Create()
	GetById()
	Update()
	DeleteById()
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r repository) GetAll() {

}

func (r repository) GetById() {

}

func (r repository) Create() {

}

func (r repository) Update() {
}

func (r repository) DeleteById() {
	// r.db.Delete()
}
