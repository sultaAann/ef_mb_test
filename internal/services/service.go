package services

import (
	"ef_md_test/internal/repositories"
	"ef_md_test/pkg/parser"
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
}

type service struct {
	r   *repositories.Repository
	psr *parser.Parser
}

func NewService(r repositories.Repository, psr *parser.Parser) Service {
	return &service{r: &r, psr: psr}
}

func (s service) GetAll() {

}
