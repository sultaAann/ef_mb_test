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

/*
1 Для получения данных с различными фильтрами и пагинацией
2 Для удаления по идентификатору
3 Для изменения сущности
4 Для добавления новых людей в формате
```json
{
"name": "Dmitriy",
"surname": "Ushakov",
"patronymic": "Vasilevich" // необязательно
}
```
*/

type Service interface {
	GetAll()
	GetById(id uint) (*models.Person, error)
	Create(dto models.RequestDTO) (uint, error)
	Update() error
	DeleteById(id uint)
}

type service struct {
	r   repositories.Repository
	psr *parser.Parser
}

func NewService(r repositories.Repository, psr *parser.Parser) Service {
	return &service{r: r, psr: psr}
}

func (s service) GetAll() {

}

func (s service) GetById(id uint) (*models.Person, error) {
	person, err := s.r.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = custom_errors.NewNotFoundError("Person", id, "Person Not Found")
			return nil, err
		}
		return nil, err
	}
	return person, nil
}

func (s service) Create(dto models.RequestDTO) (uint, error) {
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

func (s service) Update() error {
	return nil
}

func (s service) DeleteById(id uint) {
	s.r.DeleteById(id)
}
