package services

import (
	"ef_md_test/internal/custom_errors"
	"ef_md_test/internal/models"
	"ef_md_test/internal/repositories"
	"ef_md_test/pkg/parser"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Service interface {
	GetAll(page int) ([]models.Person, error)
	GetById(id uint) (*models.Person, error)
	Create(dto models.CreateDTO) (uint, error)
	Update(dto models.UpdateDTO) error
	DeleteById(id uint) error
}

type service struct {
	r   repositories.Repository
	psr *parser.Parser
}

func NewService(r repositories.Repository, psr *parser.Parser) Service {
	return &service{r: r, psr: psr}
}

func (s service) GetAll(page int) ([]models.Person, error) {
	return nil, nil
	// panic("DD")
}

func (s service) GetById(id uint) (*models.Person, error) {
	person, err := s.r.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("service")
			err = custom_errors.NewNotFoundError("Person", id, "Person Not Found")
			return nil, err
		}
		return nil, err
	}
	return person, nil
}

func (s service) Create(dto models.CreateDTO) (uint, error) {
	parsed_data, err := s.psr.Parse(dto.Name)
	if err != nil {
		return 0, fmt.Errorf("Error Parsing data: %w", err)
	}
	person := models.Person{
		Name:        dto.Name,
		Surname:     dto.Surname,
		Pantronymic: dto.Pantronymic,
		Age:         parsed_data.Age,
		Gender:      parsed_data.Gender,
		Nations:     parsed_data.Nations,
	}
	id, err := s.r.Create(person)
	if err != nil {
		return 0, fmt.Errorf("Error creating data: %w", err)
	}
	return id, nil
}

func (s service) Update(dto models.UpdateDTO) error {
	err := s.r.Update(models.Person{
		ID:          dto.ID,
		Name:        dto.Name,
		Surname:     dto.Surname,
		Pantronymic: dto.Pantronymic,
		Age:         dto.Age,
		Gender:      dto.Gender,
		Nations:     dto.Nations,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_errors.NewNotFoundError("Person", dto.ID, "Person Not Found")
		}
		return err
	}
	return nil
}

func (s service) DeleteById(id uint) error {
	return s.r.DeleteById(id)
}
