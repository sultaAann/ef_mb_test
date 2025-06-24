package repositories

import (
	"ef_md_test/internal/custom_errors"
	"ef_md_test/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll(pageSize, offset int) ([]models.Person, int64, error)
	GetById(id uint) (*models.Person, error)
	Create(person models.Person) (uint, error)
	Update(person models.Person) error
	DeleteById(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(pageSize, offset int) ([]models.Person, int64, error) {
	var people []models.Person
	result := r.db.Limit(pageSize).Offset(offset).Find(&people)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64
	r.db.Model(&models.Person{}).Count(&total)
	return people, total, nil
}

func (r *repository) GetById(id uint) (*models.Person, error) {
	person := models.Person{}
	result := r.db.First(&person, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("repository")
			fmt.Println("repository")
			return nil, err
		}
		return nil, fmt.Errorf("Failed to fetch user: %d %w", id, err)
	}
	return &person, nil
}

func (r *repository) Create(person models.Person) (uint, error) {
	result := r.db.Create(&person)
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

// не удаляет полностью делает нулл и сдавит время deleted_at. что то типо отметки. отмечает что запись удалена. но оставлена в бд
func (r *repository) DeleteById(id uint) error {
	res := r.db.Delete(&models.Person{ID: id})

	if res.Error != nil {
		return fmt.Errorf("Db Error: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return custom_errors.NewNotFoundError("", id, "Deleted data is not found")
	}
	return nil
}
