package repositories

import (
	"ef_md_test/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll()
	GetById(id uint) (*models.Person, error)
	Create(person models.Person) (uint, error)
	Update(person models.Person) error
	DeleteById(id uint)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() {

}

func (r *repository) GetById(id uint) (*models.Person, error) {
	person := models.Person{}
	result := r.db.First(&person, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("Failed to fetch user: %d %w", id, err)
	}
	return &person, nil
}

func (r *repository) Create(person models.Person) (uint, error) {
	result := r.db.Create(&person)
	fmt.Println(person)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return 0, err
		}
		return 0, fmt.Errorf("Error creating data: %w", err)
	}
	return person.ID, nil
}

func (r *repository) Update(person models.Person) error {
	var now models.Person
	if err := r.db.First(&now, person.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return fmt.Errorf("Error updating data: %w", err)
	}
	err := r.db.Model(&now).Updates(&models.Person{
		Name:        person.Name,
		Surname:     person.Surname,
		Pantronymic: person.Pantronymic,
		Age:         person.Age,
		Gender:      person.Gender,
		Nations:     person.Nations,
	}).Error
	if err != nil {
		return fmt.Errorf("Error Updating data: %w", err)
	}
	return nil
}

func (r *repository) DeleteById(id uint) {
	r.db.Delete(&models.Person{}, id)
}
